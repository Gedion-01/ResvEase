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
		Rating      int
		Amenities   []string
	}{
		{
			Name:        "Hotel Boss",
			Description: "Welcome to Hotel Boss, a unique experience in relaxation and comfort.",
			Images:      []string{"https://ak-d.tripcdn.com/images/02038120008alhyq433DF_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/220d11000000q73wq6391_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/02069120008alhhha57E1_R_600_600_R5_D.jpg_.webp"},
			Location:    "Singapore",
			Rating:      4,
			Amenities:   []string{"Free Wi-Fi", "Swimming Pool", "Spa", "Gym"},
		},
		{
			Name:        "Mountain Retreat Lodge",
			Description: "Welcome to Mountain Retreat Lodge, a unique experience in relaxation and comfort.",
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc5z12000faxcqn4B7EE_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/1mc3y12000dcojthuEBF2_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/0222i12000b3k10xw1EF0_R_600_600_R5_D.jpg_.webp"},
			Location:    "Paris",
			Rating:      3,
			Amenities:   []string{"Free Parking", "Breakfast Included", "Pet Friendly"},
		},
		{
			Name:        "Catalonia Granada",
			Description: "With a stay at Catalonia Granada Hotel, you'll be centrally located in Granada, within a 5-minute drive of Granada Cathedral and Alhambra. This spa hotel is 0.4 mi (0.7 km) from University of Granada Hospital and 0.4 mi (0.7 km) from Triunfo Gardens. amper yourself with a visit to the spa, which offers massages and body treatments. You can take advantage of recreational amenities such as an outdoor pool, a sauna, and a fitness center. Additional features at this hotel include complimentary wireless internet access, concierge services, and a banquet hall.",
			Images:      []string{"https://ak-d.tripcdn.com/images/0226f12000a6prm7821C0_R_960_660_R5_D.jpg", "https://ak-d.tripcdn.com/images/02245120009h1qd0d3626_R_339_206_R5_D.jpg", "https://ak-d.tripcdn.com/images/0226h12000a6proxp43A7_R_339_206_R5_D.jpg"},
			Location:    "Madrid",
			Rating:      4,
			Amenities:   []string{"Rooftop Bar", "Outdoor swimming pool", "Sauna", "24/7 Room Service", "Gym"},
		},
		{
			Name:        "HotelF1 Paris Porte de Châtillon",
			Description: "Welcome to Royal Heritage Inn, a unique experience in relaxation and comfort.",
			Images:      []string{"https://ak-d.tripcdn.com/images/2203180000014k0pmFD28_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/220w0u000000jh7x17BFB_R_600_600_R5_D.jpg_.webp", "https://ak-d.tripcdn.com/images/0223c12000bh55rfvBFC0_R_600_600_R5_D.jpg_.webp"},
			Location:    "Paris",
			Rating:      2,
			Amenities:   []string{"Kids' Play Area", "Airport Shuttle", "Laundry Service"},
		},
		{
			Name:        "Grand Luxor Costa Blanca",
			Description: "Welcome to Grand Luxor Costa Blanca, a unique experience in relaxation and comfort.",
			Images:      []string{"https://ak-d.tripcdn.com/images/0224n12000gr0egmeB106_R_960_660_R5_D.jpg", "https://ak-d.tripcdn.com/images/0224512000gr0e9gy851F_R_339_206_R5_D.jpg", "https://ak-d.tripcdn.com/images/1mc6b12000h6cmgha8916_R_339_206_R5_D.jpg"},
			Location:    "Benidorm Spain",
			Rating:      5,
			Amenities:   []string{"Beach Access", "Water Sports", "Live Entertainment"},
		},
		{
			Name:        "Hilton Garden Inn Bali Ngurah Rai Airport",
			Description: "A stay at Hilton Garden Inn Bali Ngurah Rai Airport - CHSE Certified places you in the heart of Tuban, within a 5-minute drive of Waterbom Bali and Kuta Beach. This 4-star hotel is 2.7 mi (4.3 km) from Beachwalk Shopping Center and 3.6 mi (5.9 km) from Krisna.",
			Images:      []string{"https://ak-d.tripcdn.com/images/0201u120008cocenc383B_R_960_660_R5_D.jpg", "https://ak-d.tripcdn.com/images/0202x120008165fhf29DA_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/0206p1200081642zq56AC_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F"},
			Location:    "Bali",
			Rating:      4,
			Amenities:   []string{"Free Wi-Fi", "Airport Shuttle", "Swimming Pool", "Spa", "Gym", "Massage Room"},
		},
		{
			Name:        "City of Aventus Hotel - Denpasar",
			Description: "With a stay at City of Aventus Hotel - Denpasar, you'll be centrally located in Denpasar, within a 10-minute drive of Kuta Beach and Eat Street. This spa hotel is 4.2 mi (6.7 km) from Legian Beach and 5 mi (8 km) from Seminyak Beach. Relax at the full-service spa, where you can enjoy massages. Additional features at this hotel include complimentary wireless internet access and concierge services.",
			Images:      []string{"https://ak-d.tripcdn.com/images/0222l12000d3en0222A77_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/0582712000dpnnq9iAD5B_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/0587312000doa8ahj49E6_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F"},
			Location:    "Bali",
			Rating:      4,
			Amenities:   []string{"Free Wi-Fi", "Airport Shuttle", "Swimming Pool", "Spa", "Gym", "Massage Room"},
		},
		{
			Name:        "New York Hilton Midtown",
			Description: "A stay at New York Hilton Midtown places you in the heart of New York, within a 5-minute walk of 5th Avenue and Broadway. This 4-star hotel is 0.3 mi (0.4 km) from Central Park and 0.4 mi (0.6 km) from St. Patrick's Cathedral. Take advantage of recreation opportunities such as a 24-hour health club, or other amenities including complimentary wireless Internet access and concierge services. Additional features at this hotel include gift shops/newsstands, a hair salon, and wedding services.",
			Images:      []string{"https://ak-d.tripcdn.com/images/220v0g00000082ymr630C_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/022041200091z4vt2510D_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/022651200091z4qza5A92_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F"},
			Location:    "New York",
			Rating:      4,
			Amenities:   []string{"Free Wi-Fi", "Swimming Pool", "Spa", "Gym", "Restaurant", "Bar"},
		},
		{
			Name:      "NYX Hotel Madrid by Leonardo Hotels",
			Images:    []string{"https://ak-d.tripcdn.com/images/220t14000000w1s851872_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/0222i120009ue0zjx55E6_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F", "https://ak-d.tripcdn.com/images/02212120009ue14klF190_W_1280_853_R5.webp?proc=watermark/image_trip1,l_ne,x_16,y_16,w_67,h_16;digimark/t_image,logo_tripbinary;ignoredefaultwm,1A8F"},
			Location:  "Madrid",
			Rating:    4,
			Amenities: []string{"Free Wi-Fi", "Swimming Pool", "Spa", "Gym", "Restaurant", "Bar"},
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

	hiltonGardenInnRooms := []types.Room{
		{
			Name:        "Standard King Room",
			Description: "Room with 1 king bed, city view, non-smoking",
			Price:       47,
			Capacity:    6,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0205g120008vuuaqsF0F3_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02032120008vurbp208AF_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0222b120008kwge0z7947_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Standard Twin Room",
			Description: "Room with 2 single beds, city view, non-smoking",
			Price:       47,
			Capacity:    6,
			BedType:     "2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0201j120008vuu6zmABA6_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02035120008vurpmi449D_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02032120008vurbp208AF_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Deluxe Pool View King Room",
			Description: "Room with 1 king bed, city view, non-smoking",
			Price:       70,
			Capacity:    6,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/02049120008vuuxrrE10E_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02032120008vurbp208AF_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02001120008163rgy5925_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Deluxe Pool View Twin Room",
			Description: "Room with 2 single beds, city view, non-smoking",
			Price:       70,
			Capacity:    6,
			BedType:     "2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0200x1200081656fjCB62_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0205d120008163j1n0FF6_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0205f1200081653x1446B_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Family Room",
			Description: "Room with 1 king bed, city view, non-smoking",
			Price:       94,
			Capacity:    5,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/02008120008vuqsclF213_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0206p120008vuqicrB3A4_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02032120008vurbp208AF_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "One Bedroom King Suite",
			Description: "Room with 1 king bed, city view, non-smoking",
			Price:       120,
			Capacity:    8,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Bathtub", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0206o1200081642zpD30B_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0206p1200081642zq56AC_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0202t1200081642gaF605_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Presidential Suite",
			Description: "Room with 1 king bed, city view, non-smoking",
			Price:       213,
			Capacity:    6,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Bathtub", "Air conditioning", "Private bathroom", "City view"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0205d120008164jba5D73_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02240120008kwgf508265_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0201z120008164czq6479_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5w12000azyx5q593CF_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	cityOfAventusRooms := []types.Room{
		{
			Name:        "Standard Room",
			Description: "Room with 1 double bed or 2 single beds, no windows, non-smoking",
			Price:       16,
			Capacity:    7,
			BedType:     "1 Double Bed or 2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc4s12000cf9osf0A08A_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6p12000cf9ovgg2489_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2312000cf9oo8856B4_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Premier Room",
			Description: "Room with 1 double bed, no windows, non-smoking",
			Price:       42,
			Capacity:    8,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc5612000cf9q3ek4F93_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5m12000cf9q3771909_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6f12000cf9q2oo88AB_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Premier Family Room",
			Description: "Room with 1 double bed and 1 single bed, non-smoking",
			Price:       57,
			Capacity:    25,
			BedType:     "1 Double Bed and 1 Single Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc6412000cf9pw0pB433_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc0312000cf9psy5583D_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2b12000cf9ppudC01F_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Premier City View Room",
			Description: "Room with 1 double bed, city view, non-smoking",
			Price:       67,
			Capacity:    13,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc5812000cf9p6a21DC5_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1e12000cf9p4e86DD8_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2212000cf9ozn82757_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Suite Room",
			Description: "Room with 1 double bed, no windows, non-smoking",
			Price:       67,
			Capacity:    19,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3w12000cf9pggj670B_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5q12000cf9phg91EED_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc2s12000cf9pn5j9D86_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Family Suite Room",
			Description: "Room with 2 double beds, non-smoking",
			Price:       67,
			Capacity:    13,
			BedType:     "2 Double Beds",
			Bedrooms:    1,
			Amenities:   []string{"Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc1v12000cf9qfdl8BC0_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5e12000cf9q5h69A34_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5y12000cf9qhdm59A1_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	newYorkHamiton := []types.Room{
		{
			Name:        "Urban Queen Room",
			Description: "Room with 1 queen bed, non-smoking",
			Price:       199,
			Capacity:    4,
			BedType:     "1 Queen Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0220j1200091z50qj5E5D_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0220w1200091z4r6d8352_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/20081e000001fn9b8E28D_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Urban King Room",
			Description: "Room with 1 king bed and 1 sofa bed, non-smoking",
			Price:       229,
			Capacity:    4,
			BedType:     "1 King Bed and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0226w1200091z4k664450_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0221j1200091z4ljc08BD_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0222r1200091z53qqC016_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Urban Two Double Room",
			Description: "Room with 2 double beds, non-smoking",
			Price:       229,
			Capacity:    5,
			BedType:     "2 Double Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0206d120009cup70o39D1_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0204x120009cupeal500A_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02067120009cupbqg6D27_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Urban Two Double Room-Hearing Accessible",
			Description: "Room with 2 double beds, non-smoking, hearing accessible",
			Price:       229,
			Capacity:    5,
			BedType:     "2 Double Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/02267120009q6ohp9F3A9_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0225a120009q6nax26A7F_R_600_400_R5.webp", "ihttps://ak-d.tripcdn.com/images/02211120009q6l0u28C1D_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Urban King Room-Hearing Accessible",
			Description: "Room with 1 king bed, non-smoking, hearing accessible",
			Price:       229,
			Capacity:    6,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0222s1200091z4lchFF29_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/022121200091z50y0B8B3_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0224t1200091z4jeh3E4C_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Urban King Room with Bathtub-Mobility Accessible",
			Description: "Room with 1 king bed, non-smoking, mobility accessible with bathtub",
			Price:       229,
			Capacity:    4,
			BedType:     "1 King Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/02053120009curyxhBD88_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02069120009cusa9i958B_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02061120009cus77u0B82_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},

		{
			Name:        "City King Room",
			Description: "Room with 1 king bed and 1 sofa bed, non-smoking, city view",
			Price:       239,
			Capacity:    7,
			BedType:     "1 King Bed and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/02068120009cup1izD02B_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0222s1200091z4lchFF29_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0225u1200091z4lidB389_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "City Two Double Room",
			Description: "Room with 2 double beds, non-smoking, city view",
			Price:       239,
			Capacity:    6,
			BedType:     "2 Double Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0220e1200091z4yjj9E7A_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0225u1200091z4lidB389_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/02267120009q6ohp9F3A9_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Skyline King Room",
			Description: "Room with 1 king bed and 1 sofa bed, city view, non-smoking",
			Price:       249,
			Capacity:    5,
			BedType:     "1 King Bed and 1 Sofa Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/0222s1200091z4lchFF29_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0223k1200091z4jm6A479_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/0225u1200091z4lidB389_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
	}

	nyxMadridRooms := []types.Room{
		{
			Name:        "Economy Double Room",
			Description: "Room with 1 double bed, non-smoking",
			Price:       95,
			Capacity:    5,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Bathtub", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc6m12000br64d8xDA75_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1d12000br6498770FD_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1d12000br645ltA38F_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Comfort Double Or Twin Room",
			Description: "Room with 1 double bed or 2 single beds, non-smoking",
			Price:       100,
			Capacity:    5,
			BedType:     "1 Double Bed or 2 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3x12000br656553F76_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc3112000br656536F59_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5n12000br64oq413B1_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Comfort Triple Room",
			Description: "Room with 3 single beds, non-smoking",
			Price:       100,
			Capacity:    5,
			BedType:     "3 Single Beds",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc3x12000br656553F76_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc6s12000br66afr6725_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1n12000br664jlEA01_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
		},
		{
			Name:        "Superior Double Room",
			Description: "Room with 1 double bed, non-smoking",
			Price:       110,
			Capacity:    4,
			BedType:     "1 Double Bed",
			Bedrooms:    1,
			Amenities:   []string{"Free Wi-Fi", "Air conditioning", "Private bathroom"},
			Images:      []string{"https://ak-d.tripcdn.com/images/1mc0b12000br65qskE635_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc5e12000br6561h22F9_R_600_400_R5.webp", "https://ak-d.tripcdn.com/images/1mc1412000br65u7w20AB_R_600_400_R5.webp"},
			Features:    []string{"Non-smoking", "Instant Confirmation", "Prepay online"},
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
		case "Hilton Garden Inn Bali Ngurah Rai Airport":
			rooms = hiltonGardenInnRooms
		case "City of Aventus Hotel - Denpasar":
			rooms = cityOfAventusRooms
		case "New York Hilton Midtown":
			rooms = newYorkHamiton
		case "NYX Hotel Madrid by Leonardo Hotels":
			rooms = nyxMadridRooms
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
