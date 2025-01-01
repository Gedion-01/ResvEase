package main

import (
	"context"
	"fmt"
	"hotel-reservation/api"
	"hotel-reservation/db"
	"hotel-reservation/db/fixtures"
	"hotel-reservation/types"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		ctx           = context.Background()
		mongoEndPoint = os.Getenv("MONGO_DB_URL")
		mongoDBName   = os.Getenv("MONGO_DB_NAME")
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndPoint))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	bookingStore := db.NewMongoBookingStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore, bookingStore),
		Hotel:   hotelStore,
	}
	user := fixtures.AddUser(store, "james", "foo", false)
	fmt.Println("james ->", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", api.CreateTokenFromUser(admin))

	var roomTypes = []types.Room{
		{
			Name:        "Deluxe Room",
			Description: "A spacious room with a beautiful city view.",
			Price:       200,
			Capacity:    2,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included", "Room Service"},
			Images: []string{
				"https://source.unsplash.com/800x600/?room",
				"https://source.unsplash.com/800x600/?bedroom",
				"https://source.unsplash.com/800x600/?city-view",
			},
			BedType:  "King",
			Bedrooms: 1,
		},
		{
			Name:        "Executive Suite",
			Description: "A luxurious suite with a separate living area.",
			Price:       350,
			Capacity:    3,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included", "Mini Bar", "Room Service"},
			Images: []string{
				"https://source.unsplash.com/800x600/?suite",
				"https://source.unsplash.com/800x600/?luxury-room",
				"https://source.unsplash.com/800x600/?living-area",
			},
			BedType:  "Queen",
			Bedrooms: 2,
		},
		{
			Name:        "Family Room",
			Description: "Perfect for families with children, includes a kids area.",
			Price:       400,
			Capacity:    4,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included", "Room Service", "Kids Area"},
			Images: []string{
				"https://source.unsplash.com/800x600/?family-room",
				"https://source.unsplash.com/800x600/?kids-area",
				"https://source.unsplash.com/800x600/?family",
			},
			BedType:  "Double",
			Bedrooms: 3,
		},
		{
			Name:        "Standard Room",
			Description: "A cozy room with basic amenities.",
			Price:       150,
			Capacity:    2,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included"},
			Images: []string{
				"https://source.unsplash.com/800x600/?standard-room",
				"https://source.unsplash.com/800x600/?cozy-room",
				"https://source.unsplash.com/800x600/?basic-room",
			},
			BedType:  "Twin",
			Bedrooms: 1,
		},
	}

	hotel := fixtures.AddHotel(
		store,
		"Luxury Resort & Spa",
		"Experience luxury and relaxation at our beachfront resort with stunning ocean views and world-class amenities.",
		[]string{
			"https://source.unsplash.com/800x600/?hotel",
			"https://source.unsplash.com/800x600/?resort",
			"https://source.unsplash.com/800x600/?beachfront",
		},
		"Maldives",
		4.8,
		[]string{"Free Wi-Fi", "Swimming Pool", "Spa", "Fitness Center", "Beach Access"},
		nil,
	)

	var addedRooms []types.Room
	for _, roomType := range roomTypes {
		for i := 0; i < 4; i++ {
			room := fixtures.AddRoom(
				store,
				roomType.Name,
				roomType.Description,
				roomType.Price,
				roomType.Capacity,
				roomType.Amenities,
				roomType.Images,
				roomType.BedType,
				roomType.Bedrooms,
				true,
				hotel.ID,
			)
			addedRooms = append(addedRooms, *room)
		}
	}

	booking := fixtures.AddBooking(store, user.ID, addedRooms[0].ID, time.Now(), time.Now().AddDate(0, 0, 2))
	fmt.Println("booking ->", booking.ID)

	// for i := 0; i < 100; i++ {
	// 	name := fmt.Sprintf("hotel name %d", i)
	// 	location := fmt.Sprintf("location %d", i)
	// 	fixtures.AddHotel(store, name, location, rand.Intn(5)+1, nil)
	// }
}
