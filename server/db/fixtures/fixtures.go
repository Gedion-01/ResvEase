package fixtures

import (
	"context"
	"fmt"
	"hotel-reservation/db"
	"hotel-reservation/types"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBooking(store *db.Store, uid, rid primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   uid,
		RoomID:   rid,
		FromDate: from,
		TillDate: till,
	}
	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}

func AddRoom(store *db.Store, name string, description string, price float64, capacity int, amenities []string, images []string, bedtype string, bedrooms int, features []string, available bool, hid primitive.ObjectID) *types.Room {
	room := &types.Room{
		Name:        name,
		Description: description,
		Price:       price,
		Capacity:    capacity,
		Amenities:   amenities,
		Images:      images,
		BedType:     bedtype,
		Bedrooms:    bedrooms,
		Features:    features,
		Available:   available,
		HotelID:     hid,
	}
	insertedRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal((err))
	}
	return insertedRoom
}

func AddHotel(store *db.Store, name string, description string, images []string, location string, rating int, amenities []string, rooms []primitive.ObjectID) *types.Hotel {
	var roomIDS = rooms
	if rooms == nil {
		roomIDS = []primitive.ObjectID{}
	}
	hotel := types.Hotel{
		Name:        name,
		Images:      images,
		Location:    location,
		Description: description,
		Amenities:   amenities,
		Rating:      rating,
		Rooms:       roomIDS,
	}
	insertedHotel, err := store.Hotel.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddUser(store *db.Store, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}
