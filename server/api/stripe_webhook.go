package api

import (
	"context"
	"encoding/json"
	"hotel-reservation/db"
	"hotel-reservation/types"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
	"go.mongodb.org/mongo-driver/bson"
)

type StripeWebhookHandler struct {
	Store *db.Store
}

func NewStripeWebhookHandler(store *db.Store) *StripeWebhookHandler {
	return &StripeWebhookHandler{Store: store}
}

func (h *StripeWebhookHandler) HandleStripeWebhook(c *fiber.Ctx) error {
	payload := c.Body()

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if endpointSecret == "" {
		log.Println("STRIPE_WEBHOOK_SECRET is not set")
		return c.Status(http.StatusInternalServerError).SendString("Webhook secret not configured")
	}

	sigHeader := c.Get("Stripe-Signature")

	event, err := webhook.ConstructEvent(payload, sigHeader, endpointSecret)
	if err != nil {
		log.Printf("Error verifying webhook signature: %v", err)
		return c.Status(http.StatusBadRequest).SendString("Signature verification failed")
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			return c.Status(http.StatusBadRequest).SendString("Failed to parse webhook JSON")
		}
		bookingID := session.Metadata["booking_id"]
		if bookingID != "" {
			err := h.Store.Booking.UpdateBooking(context.Background(), bookingID, bson.M{"status": types.Confirmed})
			if err != nil {
				log.Printf("Failed to update booking status: %v", err)
			} else {
				log.Printf("Booking %s confirmed", bookingID)
			}
		}
	case "checkout.session.expired", "payment_intent.payment_failed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			return c.Status(http.StatusBadRequest).SendString("Failed to parse webhook JSON")
		}
		bookingID := session.Metadata["booking_id"]
		if bookingID != "" {
			err := h.Store.Booking.UpdateBooking(context.Background(), bookingID, bson.M{"status": types.Canceled})
			if err != nil {
				log.Printf("Failed to update booking status: %v", err)
			} else {
				log.Printf("Booking %s canceled", bookingID)
			}
		}
	default:
		log.Printf("Unhandled event type: %s", event.Type)
	}

	return c.SendStatus(http.StatusOK)
}
