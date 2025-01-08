package main

import (
	"fmt"
	"hotel-reservation/db"
	"hotel-reservation/db/fixtures"
	"hotel-reservation/types"
	"math"
	"time"
)

func GenerateSpecificHotelsAndRooms(store *db.Store, user *types.User) {
	hotels := []struct {
		Name            string
		Description     string
		Images          []string
		Location        string
		Rating          int
		Amenities       []string
		PriceMultiplier float64 // Each hotel gets a unique price multiplier
	}{
		{
			Name:            "Sunset Paradise Hotel",
			Description:     "Welcome to Sunset Paradise Hotel, a unique experience in relaxation and comfort.",
			Images:          []string{"https://source.unsplash.com/800x600/?hotel", "https://source.unsplash.com/800x600/?resort", "https://source.unsplash.com/800x600/?luxury-hotel"},
			Location:        "Location 1",
			Rating:          2,
			Amenities:       []string{"Free Wi-Fi", "Swimming Pool", "Spa", "Gym"},
			PriceMultiplier: 1.0,
		},
		{
			Name:            "Mountain Retreat Lodge",
			Description:     "Welcome to Mountain Retreat Lodge, a unique experience in relaxation and comfort.",
			Images:          []string{"https://source.unsplash.com/800x600/?mountain", "https://source.unsplash.com/800x600/?lodge", "https://source.unsplash.com/800x600/?nature"},
			Location:        "Location 2",
			Rating:          3,
			Amenities:       []string{"Free Parking", "Breakfast Included", "Pet Friendly"},
			PriceMultiplier: 1.2,
		},
		{
			Name:            "Urban Chic Hotel",
			Description:     "Welcome to Urban Chic Hotel, a unique experience in relaxation and comfort.",
			Images:          []string{"https://source.unsplash.com/800x600/?urban", "https://source.unsplash.com/800x600/?chic", "https://source.unsplash.com/800x600/?cityscape"},
			Location:        "Location 3",
			Rating:          4,
			Amenities:       []string{"Rooftop Bar", "24/7 Room Service", "Business Center", "Gym"},
			PriceMultiplier: 1.5,
		},
		{
			Name:            "Royal Heritage Inn",
			Description:     "Welcome to Royal Heritage Inn, a unique experience in relaxation and comfort.",
			Images:          []string{"https://source.unsplash.com/800x600/?royal", "https://source.unsplash.com/800x600/?heritage", "https://source.unsplash.com/800x600/?inn"},
			Location:        "Location 4",
			Rating:          2,
			Amenities:       []string{"Kids' Play Area", "Airport Shuttle", "Laundry Service"},
			PriceMultiplier: 1.3,
		},
		{
			Name:            "Seaside Escape Resort",
			Description:     "Welcome to Seaside Escape Resort, a unique experience in relaxation and comfort.",
			Images:          []string{"https://source.unsplash.com/800x600/?seaside", "https://source.unsplash.com/800x600/?escape", "https://source.unsplash.com/800x600/?beachfront"},
			Location:        "Location 5",
			Rating:          5,
			Amenities:       []string{"Beach Access", "Water Sports", "Live Entertainment"},
			PriceMultiplier: 1.4,
		},
	}

	roomTypes := []types.Room{
		{
			Name:        "Deluxe King Room",
			Description: "Spacious room with city view",
			Price:       120,
			Capacity:    2,
			BedType:     "King",
			Bedrooms:    2,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included", "Room Service"},
			Images:      []string{"/room1-1.jpg", "/room1-2.jpg", "/room1-3.jpg"},
			Features:    []string{"Breakfast Included", "Free Cancellation", "1 Double Bed"},
		},
		{
			Name:        "Executive Queen Suite",
			Description: "Luxury suite with separate living area",
			Price:       350,
			Capacity:    2,
			BedType:     "Queen",
			Bedrooms:    2,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included", "Room Service", "Mini Bar"},
			Images:      []string{"/room2-1.jpg", "/room2-2.jpg", "/room2-3.jpg"},
			Features:    []string{"Prepay Online", "Instant Confirmation", "2 Beds"},
		},
		{
			Name:        "Family Twin Room",
			Description: "Perfect for families with children",
			Price:       400,
			Capacity:    4,
			BedType:     "Twin",
			Bedrooms:    4,
			Amenities:   []string{"Free Wi-Fi", "Breakfast Included", "Room Service", "Kids Area"},
			Images:      []string{"/room3-1.jpg", "/room3-2.jpg", "/room3-3.jpg"},
			Features:    []string{"Breakfast Included", "Instant Confirmation", "2 Beds"},
		},
		{
			Name:        "Standard Single Room",
			Description: "Cozy room for solo travelers",
			Price:       80,
			Capacity:    1,
			BedType:     "Single",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Room Service"},
			Images:      []string{"/room4-1.jpg", "/room4-2.jpg", "/room4-3.jpg"},
			Features:    []string{"Free Cancellation", "1 Single Bed"},
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

		for _, roomType := range roomTypes {
			adjustedPrice := math.Round(roomType.Price*hotelData.PriceMultiplier*100) / 100
			for i := 0; i < 4; i++ {
				addedRoom := fixtures.AddRoom(
					store,
					roomType.Name,
					roomType.Description,
					adjustedPrice,
					roomType.Capacity,
					roomType.Amenities,
					roomType.Images,
					roomType.BedType,
					roomType.Bedrooms,
					roomType.Features,
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
