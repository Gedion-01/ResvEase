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
