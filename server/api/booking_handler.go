package api

import (
	"hotel-reservation/db"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	if err := h.store.Booking.UpdateBooking(c.Context(), c.Params("id"), bson.M{"canceled": true}); err != nil {
		return err
	}
	return c.JSON(genericResp{Type: "msg", Msg: "updated"})
}

type UserBookingsParam struct {
	db.Pagination
}

func (h *BookingHandler) HandleGetUserBookings(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewError(http.StatusBadRequest, "invalid id")
	}
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid page",
		})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid limit",
		})
	}

	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{"userID": oid}, page, limit)
	if err != nil {
		return ErrResourceNotFound("bookings")
	}
	resp := ResourceResp{
		Data:    bookings,
		Results: len(bookings),
		Page:    page,
	}
	return c.JSON(resp)
}

func (h *BookingHandler) HandleGetUserBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	return c.JSON(booking)
}

// TODO: this needs to be user authorized
func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	return c.JSON(booking)
}
