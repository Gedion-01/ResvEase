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

var amenityMapping = map[string]string{
	"breakfast":    "Breakfast Included",
	"cancellation": "Free Cancellation",
	"double-bed":   "1 Double Bed",
	"two-beds":     "2 Beds",
	"prepay":       "Prepay Online",
	"instant":      "Instant Confirmation",
	"free-wifi":    "Free Wi-Fi",
	"room-service": "Room Service",
}

func translateAmenities(amenities string) []string {
	amenityList := strings.Split(amenities, ",")
	translatedAmenities := make([]string, 0, len(amenityList))
	for _, amenity := range amenityList {
		if dbValue, exists := amenityMapping[amenity]; exists {
			translatedAmenities = append(translatedAmenities, dbValue)
		} else {
			translatedAmenities = append(translatedAmenities, amenity)
		}
	}
	return translatedAmenities
}

type HotelRoomParams struct {
	db.Pagination
	Rating         int     `json:"rating"`
	HotelRating    string  `json:"hotelRating"`
	HotelAmenities string  `json:"hotelAmenities"`
	RoomCapacity   string  `json:"roomCapacity"`
	RoomAmenities  string  `json:"roomAmenities"`
	RoomBedType    string  `json:"roomBedType"`
	RoomBedrooms   string  `json:"roomBedrooms"`
	CheckIn        string  `json:"checkIn"`
	CheckOut       string  `json:"checkOut"`
	MinPrice       float64 `json:"minPrice"`
	MaxPrice       float64 `json:"maxPrice"`
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	var params HotelRoomParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}

	fromDateStr := params.CheckIn
	tillDateStr := params.CheckOut

	if fromDateStr == "" || tillDateStr == "" {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "checkIn and checkOut are required",
		})
	}

	const dateFormat = "2006-01-02"
	fromDate, err := time.Parse(dateFormat, fromDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid fromDate",
		})
	}
	fromDate = fromDate.UTC()

	tillDate, err := time.Parse(dateFormat, tillDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid tillDate",
		})
	}
	tillDate = tillDate.UTC()

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
		rating, err := strconv.Atoi(params.HotelRating)
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
	// if params.RoomAmenities != "" {
	// 	roomFilters["amenities"] = bson.M{"$all": strings.Split(params.RoomAmenities, ",")}
	// }
	if params.RoomAmenities != "" {
		translatedAmenities := translateAmenities(params.RoomAmenities)
		roomFilters["amenities"] = bson.M{"$all": translatedAmenities}
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
	if params.MinPrice > 0 {
		roomFilters["price"] = bson.M{"$gte": params.MinPrice}
	}
	if params.MaxPrice > 0 {
		if _, ok := roomFilters["price"]; ok {
			roomFilters["price"].(bson.M)["$lte"] = params.MaxPrice
		} else {
			roomFilters["price"] = bson.M{"$lte": params.MaxPrice}
		}
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
	Rating    int     `json:"rating"`
	Location  string  `json:"location"`
	Amenities string  `json:"amenities"`
	MinPrice  float64 `json:"minPrice"`
	MaxPrice  float64 `json:"maxPrice"`
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
	if params.Amenities != "" {
		filter["amenities"] = bson.M{"$all": strings.Split(params.Amenities, ",")}
	}
	if params.Location != "" {
		filter["location"] = bson.M{"$regex": primitive.Regex{Pattern: params.Location, Options: "i"}}
	}

	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination, params.MinPrice, params.MaxPrice)

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
