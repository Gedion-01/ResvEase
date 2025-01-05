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

	roomBasePrices := map[string]float64{
		"Deluxe Room":     200,
		"Executive Suite": 350,
		"Standard Room":   150,
		"Family Room":     400,
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

		for roomName, basePrice := range roomBasePrices {
			adjustedPrice := basePrice * hotelData.PriceMultiplier
			for i := 0; i < 4; i++ {
				addedRoom := fixtures.AddRoom(
					store,
					roomName,
					fmt.Sprintf("A %s in %s.", roomName, hotelData.Name),
					adjustedPrice,
					2,
					[]string{"Free Wi-Fi", "Air Conditioning"},
					[]string{"https://source.unsplash.com/800x600/?" + roomName, "https://source.unsplash.com/800x600/?room"},
					"King",
					1,
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
