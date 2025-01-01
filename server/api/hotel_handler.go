package api

import (
	"hotel-reservation/db"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

type HotelRoomParams struct {
	db.Pagination
	Rating         float64 `json:"rating"`
	HotelRating    string  `json:"hotelRating"`
	HotelAmenities string  `json:"hotelAmenities"`
	HotelLocation  string  `json:"hotelLocation"`
	RoomCapacity   string  `json:"roomCapacity"`
	RoomAmenities  string  `json:"roomAmenities"`
	RoomBedType    string  `json:"roomBedType"`
	RoomBedrooms   string  `json:"roomBedrooms"`
	FromDate       string  `json:"fromDate"`
	TillDate       string  `json:"tillDate"`
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	// filter := bson.M{"hotelID": oid}
	// rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	// if err != nil {
	// 	return ErrResourceNotFound("hotel")
	// }
	// return c.JSON(rooms)
	var params HotelRoomParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}

	fromDateStr := params.FromDate
	tillDateStr := params.TillDate

	if fromDateStr == "" || tillDateStr == "" {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "fromDate and tillDate are required",
		})
	}
	fromDate, err := time.Parse(time.RFC3339, fromDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid fromDate",
		})
	}

	tillDate, err := time.Parse(time.RFC3339, tillDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid tillDate",
		})
	}

	// Parse pagination parameters
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

	// Build hotel filters
	hotelFilters := bson.M{"hotelID": oid}
	if params.HotelRating != "" {
		rating, err := strconv.ParseFloat(params.HotelRating, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(genericResp{
				Type: "error",
				Msg:  "invalid hotelRating",
			})
		}
		hotelFilters["hotel.rating"] = bson.M{"$gte": rating}
	}
	if params.HotelAmenities != "" {
		hotelFilters["hotel.amenities"] = bson.M{"$all": strings.Split(params.HotelAmenities, ",")}
	}
	if params.HotelLocation != "" {
		hotelFilters["hotel.location"] = params.HotelLocation
	}

	// Build room filters
	roomFilters := bson.M{}
	if params.RoomCapacity != "" {
		capacity, err := strconv.Atoi(params.RoomCapacity)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(genericResp{
				Type: "error",
				Msg:  "invalid roomCapacity",
			})
		}
		roomFilters["capacity"] = bson.M{"$gte": capacity}
	}
	if params.RoomAmenities != "" {
		roomFilters["amenities"] = bson.M{"$all": strings.Split(params.RoomAmenities, ",")}
	}
	if params.RoomBedType != "" {
		roomFilters["bedType"] = params.RoomBedType
	}
	if params.RoomBedrooms != "" {
		bedrooms, err := strconv.Atoi(params.RoomBedrooms)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(genericResp{
				Type: "error",
				Msg:  "invalid roomBedrooms",
			})
		}
		roomFilters["bedrooms"] = bson.M{"$gte": bedrooms}
	}

	rooms, err := h.store.Room.GetRooms(c.Context(), hotelFilters, roomFilters, &fromDate, &tillDate, page, limit)
	if err != nil {
		return err
	}
	resp := ResourceResp{
		Data:    rooms,
		Results: len(rooms),
		Page:    int(page),
	}
	return c.JSON(resp)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating float64 `json:"rating"`
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	// filter := db.Map{
	// 	"rating": params.Rating,
	// }
	filter := db.Map{}

	if params.Rating != 0 {
		filter["rating"] = params.Rating
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrResourceNotFound("hotel")
	}
	resp := ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(params.Page),
	}
	return c.JSON(resp)
}
