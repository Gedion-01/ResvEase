package db

import (
	"context"
	"hotel-reservation/types"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRoomByID(context.Context, primitive.ObjectID) (*types.Room, error)
	GetRoomsSearch(context.Context, bson.M) ([]*types.Room, error)
	GetRooms(ctx context.Context, hotelFilters, roomFilters bson.M, fromDate, tillDate *time.Time, page, limit int) ([]*types.GroupedRoom, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	HotelStore
	BookingStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore, bookingStore BookingStore) *MongoRoomStore {
	dbname := os.Getenv(MongoDbNameEnvName)
	return &MongoRoomStore{
		client:       client,
		coll:         client.Database(dbname).Collection("rooms"),
		HotelStore:   hotelStore,
		BookingStore: bookingStore,
	}
}

func (s *MongoRoomStore) GetRooms(
	ctx context.Context,
	hotelFilters, roomFilters bson.M,
	fromDate, tillDate *time.Time,
	page, limit int,
) ([]*types.GroupedRoom, error) {
	pipeline := []bson.M{
		// Step 1: Filter rooms by hotelID
		{
			"$match": bson.M{"hotelID": hotelFilters["hotelID"]},
		},
		// Step 2: Apply room filters (e.g., price, capacity)
		{
			"$match": roomFilters,
		},
		// Step 3: Look up bookings for each room
		{
			"$lookup": bson.M{
				"from":         "bookings",
				"localField":   "_id",
				"foreignField": "roomID",
				"as":           "bookings",
			},
		},
		// Step 4: Check for overlapping bookings (INCLUSIVE comparison)
		{
			"$addFields": bson.M{
				"isBooked": bson.M{
					"$gt": []interface{}{
						bson.M{
							"$size": bson.M{
								"$filter": bson.M{
									"input": "$bookings",
									"as":    "booking",
									"cond": bson.M{
										"$and": []bson.M{
											{"$eq": []interface{}{"$$booking.canceled", false}},
											// Overlap condition: booking.fromDate <= tillDate AND booking.tillDate >= fromDate
											{"$lte": []interface{}{"$$booking.fromDate", tillDate}},
											{"$gte": []interface{}{"$$booking.tillDate", fromDate}},
										},
									},
								},
							},
						},
						0,
					},
				},
			},
		},
		// Step 5: Group by room name and hotelID
		{
			"$group": bson.M{
				"_id": bson.M{
					"hotelID": "$hotelID",
					"name":    "$name",
				},
				"id":          bson.M{"$first": "$_id"},
				"name":        bson.M{"$first": "$name"},
				"description": bson.M{"$first": "$description"},
				"price":       bson.M{"$first": "$price"},
				"capacity":    bson.M{"$first": "$capacity"},
				"amenities":   bson.M{"$first": "$amenities"},
				"images":      bson.M{"$first": "$images"},
				"bedType":     bson.M{"$first": "$bedType"},
				"bedrooms":    bson.M{"$first": "$bedrooms"},
				"totalCount":  bson.M{"$sum": 1}, // Total rooms of this type
				"bookedCount": bson.M{"$sum": bson.M{"$cond": []interface{}{"$isBooked", 1, 0}}},
			},
		},
		// Step 6: Calculate availableCount
		{
			"$addFields": bson.M{
				"availableCount": bson.M{
					"$subtract": []interface{}{"$totalCount", "$bookedCount"},
				},
			},
		},
		// Step 7: Pagination
		{
			"$skip": (page - 1) * limit,
		},
		{
			"$limit": limit,
		},
		// Step 8: Final projection
		{
			"$project": bson.M{
				"_id":            0,
				"id":             1,
				"name":           1,
				"description":    1,
				"price":          1,
				"capacity":       1,
				"amenities":      1,
				"images":         1,
				"bedType":        1,
				"bedrooms":       1,
				"availableCount": 1,
			},
		},
	}

	cursor, err := s.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*types.GroupedRoom
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MongoRoomStore) GetRoomByID(ctx context.Context, id primitive.ObjectID) (*types.Room, error) {
	resp := s.coll.FindOne(ctx, bson.M{"_id": id})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	var room types.Room
	if err := resp.Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *MongoRoomStore) GetRoomsSearch(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*types.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = resp.InsertedID.(primitive.ObjectID)

	// update the hotel with this id
	filter := Map{"_id": room.HotelID}
	update := Map{"$push": bson.M{"rooms": room.ID}}

	if err := s.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}
