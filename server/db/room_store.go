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

	pipeline := []bson.M{}

	// Match the hotel filters (contains _id if you want a specific hotel)
	pipeline = append(pipeline, bson.M{
		"$match": hotelFilters,
	})

	// Look up booking documents that reference these rooms
	pipeline = append(pipeline, bson.M{
		"$lookup": bson.M{
			"from":         "bookings",
			"localField":   "_id",
			"foreignField": "roomID",
			"as":           "bookings",
		},
	})

	// Unwind bookings so we can filter on date range
	pipeline = append(pipeline, bson.M{
		"$unwind": bson.M{
			"path":                       "$bookings",
			"preserveNullAndEmptyArrays": true,
		},
	})

	// Filter out rooms that are booked within [fromDate, tillDate] (unless canceled)
	pipeline = append(pipeline, bson.M{
		"$match": bson.M{
			"$or": []bson.M{
				// No bookings at all => keep room
				{"bookings": bson.M{"$exists": false}},
				// Booking is canceled => keep room
				{"bookings.canceled": true},
				// Booking starts after the requested end => no overlap
				{"bookings.fromDate": bson.M{"$gte": tillDate}},
				// Booking ends before the requested start => no overlap
				{"bookings.tillDate": bson.M{"$lte": fromDate}},
			},
		},
	})

	// If more room-specific filters exist, apply them
	if len(roomFilters) > 0 {
		pipeline = append(pipeline, bson.M{"$match": roomFilters})
	}

	// Group by room name to count how many individual rooms share that name
	pipeline = append(pipeline, bson.M{
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
			"availableCount": bson.M{
				"$sum": bson.M{
					"$cond": []interface{}{"$available", 1, 0},
				},
			},
		},
	})

	// Pagination
	pipeline = append(pipeline, bson.M{"$skip": (page - 1) * limit})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	// Final projection
	pipeline = append(pipeline, bson.M{
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
	})

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
