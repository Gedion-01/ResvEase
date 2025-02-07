package api

import (
	"context"
	"errors"
	"fmt"
	"hotel-reservation/db"
	"hotel-reservation/types"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	HotelID        primitive.ObjectID `json:"hotelID"`
	RoomName       string             `json:"roomName"`
	FirstName      string             `json:"firstName"`
	LastName       string             `json:"lastName"`
	Phone          string             `json:"phone"`
	SpecialRequest string             `json:"specialRequest"`
	FromDate       string             `json:"fromDate"`
	TillDate       string             `json:"tillDate"`
	NumPersons     int                `json:"numPersons"`
}

// func (p BookRoomParams) validate() error {
// 	now := time.Now()
// 	if now.After(p.FromDate) || now.After(p.TillDate) {
// 		return fmt.Errorf("can't book a room in the past")
// 	}
// 	return nil
// }

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

type RoomQueryParams struct {
	db.Pagination
}

func (h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	// rooms, err := h.store.Room.GetRooms(c.Context(), bson.M{})
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(rooms)
	return nil
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	// if err := params.validate(); err != nil {
	// 	return err
	// }

	// Parse dates from "yyyy-MM-dd" format
	fromDate, err := time.Parse("2006-01-02", params.FromDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid from date format",
		})
	}
	tillDate, err := time.Parse("2006-01-02", params.TillDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid till date format",
		})
	}

	hotelID := params.HotelID

	// Find an available room by hotel and room name
	roomIDPtr, err := h.getAvailableRoom(c.Context(), hotelID, params.RoomName, params)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  fmt.Sprintf("no available rooms: %v", err),
		})
	}
	roomID := *roomIDPtr

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), hotelID.Hex())

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "failed to get hotel",
		})
	}

	// Fetch room details to calculate the price
	room, err := h.store.Room.GetRoomByID(c.Context(), roomID)
	if err != nil {
		return err
	}

	// Calculate price: number of nights * room.Price (assumed USD) converted to cents.
	numNights := int(tillDate.Sub(fromDate).Hours() / 24)
	totalAmount := room.Price * float64(numNights) * 100

	// Create a Stripe Checkout Session
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	domain := os.Getenv("DOMAIN")
	checkoutParams := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:   stripe.String("usd"),
					UnitAmount: stripe.Int64(int64(totalAmount)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(fmt.Sprintf("%s - %s", hotel.Name, room.Name)),
						Description: stripe.String(fmt.Sprintf(
							"Stay in %s at %s\nFrom: %s\nTo: %s",
							room.Name,
							hotel.Name,
							fromDate.Format("Jan 2, 2006"),
							tillDate.Format("Jan 2, 2006"),
						)),
						Images: stripe.StringSlice([]string{
							room.Images[0], // change to valid URL
						}),
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/booking-success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(domain + "/booking-cancel"),
		Metadata: map[string]string{
			"room_id":     roomID.Hex(),
			"user_id":     user.ID.Hex(),
			"from_date":   fromDate.Format(time.RFC3339),
			"till_date":   tillDate.Format(time.RFC3339),
			"num_persons": fmt.Sprintf("%d", params.NumPersons),
		},
	}

	s, err := session.New(checkoutParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "failed to create checkout session",
		})
	}

	// Return the session URL for frontend redirection
	return c.JSON(fiber.Map{"sessionUrl": s.URL})
}

// func (h *RoomHandler) isRoomAvailabeForBooking(ctx context.Context, roomID primitive.ObjectID, params BookRoomParams) (bool, error) {
// 	where := bson.M{
// 		"roomID": roomID,
// 		"fromDate": bson.M{
// 			"$gte": params.FromDate,
// 		},
// 		"tillDate": bson.M{
// 			"$lte": params.TillDate,
// 		},
// 	}
// 	bookings, err := h.store.Booking.GetBookings(ctx, where)
// 	if err != nil {
// 		return false, err
// 	}
// 	ok := len(bookings) == 0

// 	return ok, nil
// }

func (h *RoomHandler) getAvailableRoom(ctx context.Context, hotelID primitive.ObjectID, roomName string, params BookRoomParams) (*primitive.ObjectID, error) {
	rooms, err := h.store.Room.GetRoomsSearch(ctx, bson.M{
		"hotelID": hotelID,
		"name":    roomName,
	})

	if err != nil {
		return nil, err
	}

	if len(rooms) == 0 {
		return nil, errors.New("no rooms found with the given name in this hotel")
	}

	for _, room := range rooms {
		where := bson.M{
			"roomID": room.ID,
			"fromDate": bson.M{
				"$gte": params.FromDate,
			},
			"tillDate": bson.M{
				"$lte": params.TillDate,
			},
		}

		bookings, err := h.store.Booking.GetBookings(ctx, where)
		if err != nil {
			return nil, err
		}
		if len(bookings) == 0 {
			return &room.ID, nil
		}
	}
	return nil, errors.New("no available rooms for the given date range in this hotel")
}
