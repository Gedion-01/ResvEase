package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string               `bson:"name" json:"name"`
	Images      []string             `bson:"images" json:"images"`
	Location    string               `bson:"location" json:"location"`
	Description string               `bson:"description" json:"description"`
	Rating      float64              `bson:"rating" json:"rating"`
	Amenities   []string             `bson:"amenities" json:"amenities"`
	Rooms       []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Prices      []float64            `bson:"prices" json:"prices"`
}

type Room struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Images      []string           `bson:"images" json:"images"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Capacity    int                `bson:"capacity" json:"capacity"`
	Amenities   []string           `bson:"amenities" json:"amenities"`
	BedType     string             `bson:"bedType" json:"bedType"`
	Bedrooms    int                `bson:"bedrooms" json:"bedrooms"`
	Available   bool               `bson:"available" json:"available"`
	HotelID     primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}

type GroupedRoom struct {
	ID             primitive.ObjectID `json:"id"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	Price          int                `json:"price"`
	Capacity       int                `json:"capacity"`
	BedType        string             `json:"bedType"`
	Bedrooms       int                `json:"bedrooms"`
	AvailableCount int                `json:"availableCount"`
	Amenities      []string           `json:"amenities"`
	Images         []string           `json:"images"`
}
