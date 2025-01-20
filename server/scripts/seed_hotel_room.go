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
			Name:            "Hotel Boss",
			Description:     "Welcome to Hotel Boss, a unique experience in relaxation and comfort.",
			Images:          []string{"https://ak-d.tripcdn.com/images/02038120008alhyq433DF_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/220d11000000q73wq6391_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/02069120008alhhha57E1_R_600_600_R5_D.jpg_.webp"},
			Location:        "Singapore",
			Rating:          4,
			Amenities:       []string{"Free Wi-Fi", "Swimming Pool", "Spa", "Gym"},
			PriceMultiplier: 1.0,
		},
		{
			Name:            "Mountain Retreat Lodge",
			Description:     "Welcome to Mountain Retreat Lodge, a unique experience in relaxation and comfort.",
			Images:          []string{"https://ak-d.tripcdn.com/images/1mc5z12000faxcqn4B7EE_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/1mc3y12000dcojthuEBF2_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/0222i12000b3k10xw1EF0_R_600_600_R5_D.jpg_.webp"},
			Location:        "Paris",
			Rating:          3,
			Amenities:       []string{"Free Parking", "Breakfast Included", "Pet Friendly"},
			PriceMultiplier: 1.2,
		},
		{
			Name:            "Catalonia Granada",
			Description:     "With a stay at Catalonia Granada Hotel, you'll be centrally located in Granada, within a 5-minute drive of Granada Cathedral and Alhambra. This spa hotel is 0.4 mi (0.7 km) from University of Granada Hospital and 0.4 mi (0.7 km) from Triunfo Gardens. amper yourself with a visit to the spa, which offers massages and body treatments. You can take advantage of recreational amenities such as an outdoor pool, a sauna, and a fitness center. Additional features at this hotel include complimentary wireless internet access, concierge services, and a banquet hall.",
			Images:          []string{"https://ak-d.tripcdn.com/images/0226f12000a6prm7821C0_R_960_660_R5_D.jpg", "https://ak-d.tripcdn.com/images/02245120009h1qd0d3626_R_339_206_R5_D.jpg", "https://ak-d.tripcdn.com/images/0226h12000a6proxp43A7_R_339_206_R5_D.jpg"},
			Location:        "Madrid",
			Rating:          4,
			Amenities:       []string{"Rooftop Bar", "Outdoor swimming pool", "Sauna", "24/7 Room Service", "Gym"},
			PriceMultiplier: 1.5,
		},
		{
			Name:            "HotelF1 Paris Porte de Châtillon",
			Description:     "Welcome to Royal Heritage Inn, a unique experience in relaxation and comfort.",
			Images:          []string{"https://ak-d.tripcdn.com/images/2203180000014k0pmFD28_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/220w0u000000jh7x17BFB_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/0223c12000bh55rfvBFC0_R_600_600_R5_D.jpg_.webp"},
			Location:        "Paris",
			Rating:          2,
			Amenities:       []string{"Kids' Play Area", "Airport Shuttle", "Laundry Service"},
			PriceMultiplier: 1.3,
		},
		{
			Name:            "Grand Luxor Costa Blanca",
			Description:     "Welcome to Grand Luxor Costa Blanca, a unique experience in relaxation and comfort.",
			Images:          []string{"https://ak-d.tripcdn.com/images/0224n12000gr0egmeB106_R_960_660_R5_D.jpg", "https://ak-d.tripcdn.com/images/0224512000gr0e9gy851F_R_339_206_R5_D.jpg", "https://ak-d.tripcdn.com/images/1mc6b12000h6cmgha8916_R_339_206_R5_D.jpg"},
			Location:        "Benidorm Spain",
			Rating:          5,
			Amenities:       []string{"Beach Access", "Water Sports", "Live Entertainment"},
			PriceMultiplier: 1.4,
		},
	}

	hotelBoss := []types.Room{
		{
			Name:        "Superior Double Room",
			Description: "Room with 1 queen bed, non-smoking",
			Price:       109,
			Capacity:    4,
			BedType:     "1 Queen Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0205d12000818xs9l86AA_R_400_200_R5.webp", "https://ak-d.tripcdn.com/images/0202s120008alprafF09E_R_200_100_R5.webp", "https://ak-d.tripcdn.com/images/0200x12000818iwpnD0F8_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Superior Twin Room",
			Description: "Room with 2 single beds, non-smoking",
			Price:       109,
			Capacity:    4,
			BedType:     "2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc0z12000dplv8hf20A0_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5a12000dmg0wni11B5_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0200x12000818iwpnD0F8_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Superior Twin Room with City View",
			Description: "Room with 2 single beds and city view, non-smoking",
			Price:       115,
			Capacity:    4,
			BedType:     "2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0224r12000bjke3fwBA0A_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc0n12000dmmxowr53C6_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0201t12000818z8nr8EAC_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Superior Queen Room with City View",
			Description: "Room with 1 queen bed and city view, non-smoking",
			Price:       115,
			Capacity:    4,
			BedType:     "1 Queen Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc0l12000dmg1u1z86EC_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02062120008alqp9i49A3_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1c12000dplu4lv0A2B_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Superior Room With Balcony",
			Description: "Room with 1 queen bed and balcony, non-smoking",
			Price:       133,
			Capacity:    4,
			BedType:     "1 Queen Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "Balcony"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0200x12000818x8le33D7_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6s12000dmmwx8x8FFB_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6z12000dmmwyjrEF20_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Triple Room",
			Description: "Room with 1 single bed and 1 double bed, non-smoking",
			Price:       170,
			Capacity:    4,
			BedType:     "1 Single Bed and 1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc1i12000dmmzlhv8326_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2012000dmmzf6f1631_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1c12000dplu4lv0A2B_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	mountainRetreat := []types.Room{
		{
			Name:        "Standard Room With 1 Double Bed",
			Description: "Room with 1 double bed, non-smoking",
			Price:       146,
			Capacity:    4,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc0r12000eyzqmgb77B1_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0225b120008qwx152626C_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0226h120008qwwnqy17AD_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Standard Room With Twin Beds",
			Description: "Room with 2 single beds, non-smoking",
			Price:       146,
			Capacity:    6,
			BedType:     "2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc2s12000eyzp431CB60_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0222n120009c77rh7DB3D_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0223s120008qwx3pmAA1E_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Standard Room With 1 Double Bed And Views Of The Eiffel Tower",
			Description: "Room with 1 double bed and city view, non-smoking",
			Price:       146,
			Capacity:    4,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3u12000eyztgfu9D30_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc0r12000eyzqmgb77B1_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc0r12000eyzqmgb77B1_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	cataloniaGranada := []types.Room{
		{
			Name:        "Double Room",
			Description: "Room with 1 double bed or 2 single beds, non-smoking",
			Price:       50,
			Capacity:    2,
			BedType:     "1 Double Bed or 2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/200a1f000001g7dnr873C_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0222o12000adrtbek7A71_R_600_400_R5.webps", "https://ak-d.tripcdn.com/images/1mc5p12000bj3mnr9EF94_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Triple Room",
			Description: "Room with 2 single beds and 1 sofa bed, city view, non-smoking",
			Price:       66,
			Capacity:    3,
			BedType:     "2 Single Beds and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0220l12000adrtm816C65_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0582n12000cunkfwz9105_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0585y12000cungz34A864_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Superior Room",
			Description: "Room with 1 king bed, non-smoking",
			Price:       75,
			Capacity:    2,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0221m120009zrgvaa0BBB_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0580l12000cunknqyB111_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0581p12000cunkl61DAC7_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Junior Suite",
			Description: "Room with 1 king bed, non-smoking",
			Price:       100,
			Capacity:    2,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0580z12000cunkf6m8ADD_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0223f12000adrti8h24D6_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0587412000cunku0g6301_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	hotelF1Rooms := []types.Room{
		{
			Name:        "Side Car Room",
			Description: "Comfortable room with 2 single beds",
			Price:       49,
			Capacity:    2,
			BedType:     "2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Shower"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0223c12000bh55rfvBFC0_R_400_200_R5.webp", "https://ak-d.tripcdn.com/images/0220u12000bi331fg42EF_R_200_100_R5.webp", "https://ak-d.tripcdn.com/images/0221c12000bi330f6F0B9_R_200_100_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Break Room",
			Description: "Spacious room with 3 single beds",
			Price:       59,
			Capacity:    3,
			BedType:     "3 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc4012000ecpntadCAC4_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0221812000dsponrm7DD6_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc0q12000ecpnq44AE51_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Cabrio Room With Private Bathroom",
			Description: "Room with 1 double bed and private bathroom",
			Price:       64,
			Capacity:    2,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc5a12000ecppmjxFBB5_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc3c12000ecpq8vv5C8D_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0223n120009lr7s3289E3_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Cabrio Family Room With Private Bathroom",
			Description: "Family room with 1 double bed and 2 single beds",
			Price:       103,
			Capacity:    4,
			BedType:     "1 Double Bed and 2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0220o120009tjtjg24682_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1412000ecpr77f6DF1_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6012000ecpro8p57CB_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	grandLuxorRooms := []types.Room{
		{
			Name:        "Guest Room, Double, Sofa Bed, Mountain View, Balcony",
			Description: "Room with 1 double bed and 1 sofa bed, mountain view, balcony, non-smoking",
			Price:       84,
			Capacity:    4,
			BedType:     "1 Double Bed and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "Mountain view", "Balcony"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3i12000h6cjeqv2CB8_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc3d12000h6ci2gi877E_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6b12000h6cia7t9FF5_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Pay at hotel"},
		},
		{
			Name:        "Guest Room, 2 Twin, Sofa Bed, Mountain View, Balcony",
			Description: "Room with 2 single beds and 1 sofa bed, mountain view, balcony, non-smoking",
			Price:       84,
			Capacity:    4,
			BedType:     "2 Single Beds and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "Mountain view", "Balcony"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3i12000h6cjeqv2CB8_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc3d12000h6ci2gi877E_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6b12000h6cia7t9FF5_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Pay at hotel"},
		},
		{
			Name:        "Classic, Guest Room, 2 Twin, Sofa Bed, Sea View, Balcony",
			Description: "Room with 2 single beds and 1 sofa bed, ocean view, balcony, non-smoking",
			Price:       108,
			Capacity:    6,
			BedType:     "2 Single Beds and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "Ocean view", "Balcony"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc6b12000h6cmgha8916_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2p12000h6cm7zm8E04_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5y12000h6cmbqb9B9F_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc4w12000h6chlxe8FFC_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Pay at hotel"},
		},
		{
			Name:        "Classic, Guest Room, 1 Double, Sofa Bed, Sea View, Balcony",
			Description: "Room with 1 double bed and 1 sofa bed, ocean view, balcony, non-smoking",
			Price:       108,
			Capacity:    6,
			BedType:     "1 Double Bed and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "Ocean view", "Balcony"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc6b12000h6cmgha8916_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2p12000h6cm7zm8E04_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5y12000h6cmbqb9B9F_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Pay at hotel"},
		},
		{
			Name:        "Premium, Guest Room, 1 Queen, Sofa Bed, Sea View, Balcony",
			Description: "Room with 1 queen bed and 1 sofa bed, ocean view, balcony, non-smoking",
			Price:       112,
			Capacity:    7,
			BedType:     "1 Queen Bed and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "Ocean view", "Balcony"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3m12000h6cm3s45BE0_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc4812000h6clo699DB9_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6c12000h6clmxy9A9F_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1312000h6cltirFC1C_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Pay at hotel"},
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

		var rooms []types.Room
		switch hotelData.Name {
		case "Hotel Boss":
			rooms = hotelBoss
		case "Mountain Retreat Lodge":
			rooms = mountainRetreat
		case "Catalonia Granada":
			rooms = cataloniaGranada
		case "HotelF1 Paris Porte de Châtillon":
			rooms = hotelF1Rooms
		case "Grand Luxor Costa Blanca":
			rooms = grandLuxorRooms
		}

		for _, roomData := range rooms {
			for i := 0; i < 4; i++ {
				addedRoom := fixtures.AddRoom(
					store,
					roomData.Name,
					roomData.Description,
					roomData.Price,
					roomData.Capacity,
					roomData.Amenities,
					roomData.Images,
					roomData.BedType,
					roomData.Bedrooms,
					roomData.Features,
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
