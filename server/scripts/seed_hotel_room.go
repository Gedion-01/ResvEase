package main

import (
	"fmt"
	"hotel-reservation/db"
	"hotel-reservation/db/fixtures"
	"hotel-reservation/types"
	"time"
)

func GenerateSpecificHotelsAndRooms(store *db.Store, user *types.User) {
	hotels := []struct {
		Name        string
		Description string
		Images      []string
		Location    string
		Rating      float64
		Amenities   []string
	}{
		{
			"Sunset Paradise Hotel",
			"Welcome to Sunset Paradise Hotel, a unique experience in relaxation and comfort.",
			[]string{"https://source.unsplash.com/800x600/?hotel", "https://source.unsplash.com/800x600/?resort", "https://source.unsplash.com/800x600/?luxury-hotel"},
			"Location 1",
			4.5,
			[]string{"Free Wi-Fi", "Swimming Pool", "Spa", "Gym"},
		},
		{
			"Mountain Retreat Lodge",
			"Welcome to Mountain Retreat Lodge, a unique experience in relaxation and comfort.",
			[]string{"https://source.unsplash.com/800x600/?mountain", "https://source.unsplash.com/800x600/?lodge", "https://source.unsplash.com/800x600/?nature"},
			"Location 2",
			4.7,
			[]string{"Free Parking", "Breakfast Included", "Pet Friendly"},
		},
		{
			"Urban Chic Hotel",
			"Welcome to Urban Chic Hotel, a unique experience in relaxation and comfort.",
			[]string{"https://source.unsplash.com/800x600/?urban", "https://source.unsplash.com/800x600/?chic", "https://source.unsplash.com/800x600/?cityscape"},
			"Location 3",
			4.3,
			[]string{"Rooftop Bar", "24/7 Room Service", "Business Center"},
		},
		{
			"Royal Heritage Inn",
			"Welcome to Royal Heritage Inn, a unique experience in relaxation and comfort.",
			[]string{"https://source.unsplash.com/800x600/?royal", "https://source.unsplash.com/800x600/?heritage", "https://source.unsplash.com/800x600/?inn"},
			"Location 4",
			4.8,
			[]string{"Kids' Play Area", "Airport Shuttle", "Laundry Service"},
		},
		{
			"Seaside Escape Resort",
			"Welcome to Seaside Escape Resort, a unique experience in relaxation and comfort.",
			[]string{"https://source.unsplash.com/800x600/?seaside", "https://source.unsplash.com/800x600/?escape", "https://source.unsplash.com/800x600/?beachfront"},
			"Location 5",
			4.9,
			[]string{"Beach Access", "Water Sports", "Live Entertainment"},
		},
	}
	var addedRooms []types.Room

	for _, hotelData := range hotels {
		hotel := fixtures.AddHotel(
			store,
			hotelData.Name,
			hotelData.Description,
			hotelData.Images,
			hotelData.Location,
			hotelData.Rating,
			hotelData.Amenities,
			nil,
		)

		rooms := []struct {
			Name        string
			Description string
			Price       float64
			Capacity    int
			Amenities   []string
			Images      []string
			BedType     string
			Bedrooms    int
		}{
			{
				"Deluxe Room",
				"A spacious Deluxe Room.",
				200,
				2,
				[]string{"Free Wi-Fi", "Breakfast Included"},
				[]string{"https://source.unsplash.com/800x600/?deluxe", "https://source.unsplash.com/800x600/?luxury"},
				"King",
				1,
			},
			{
				"Executive Suite",
				"A luxurious Executive Suite.",
				350,
				3,
				[]string{"Mini Bar", "Room Service", "City View"},
				[]string{"https://source.unsplash.com/800x600/?suite", "https://source.unsplash.com/800x600/?executive"},
				"Queen",
				2,
			},
			{
				"Standard Room",
				"A cozy Standard Room.",
				150,
				2,
				[]string{"Balcony", "Ocean View", "Free Parking"},
				[]string{"https://source.unsplash.com/800x600/?standard", "https://source.unsplash.com/800x600/?cozy"},
				"Twin",
				1,
			},
			{
				"Family Room",
				"A spacious Family Room.",
				400,
				4,
				[]string{"Kids Area", "Flat-screen TV", "Air Conditioning"},
				[]string{"https://source.unsplash.com/800x600/?family", "https://source.unsplash.com/800x600/?kids"},
				"Double",
				2,
			},
		}

		for _, room := range rooms {
			for i := 0; i < 4; i++ {
				addedRoom := fixtures.AddRoom(
					store,
					room.Name,
					room.Description,
					room.Price,
					room.Capacity,
					room.Amenities,
					room.Images,
					room.BedType,
					room.Bedrooms,
					true,
					hotel.ID,
				)
				addedRooms = append(addedRooms, *addedRoom)
			}
		}
	}

	booking := fixtures.AddBooking(store, user.ID, addedRooms[0].ID, time.Now(), time.Now().AddDate(0, 0, 2))
	fmt.Println("booking ->", booking.ID)
}
