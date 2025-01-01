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
	// GetRooms(context.Context, bson.M) ([]map[string]interface{}, error)
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

func (s *MongoRoomStore) GetRooms(ctx context.Context, hotelFilters, roomFilters bson.M, fromDate, tillDate *time.Time, page, limit int) ([]*types.GroupedRoom, error) {
	pipeline := []bson.M{
		{"$lookup": bson.M{
			"from":         "hotels",
			"localField":   "hotelID",
			"foreignField": "_id",
			"as":           "hotel",
		}},
		{"$unwind": "$hotel"},
		{"$match": hotelFilters},
		{"$lookup": bson.M{
			"from":         "bookings",
			"localField":   "_id",
			"foreignField": "roomID",
			"as":           "bookings",
		}},
		{"$unwind": bson.M{
			"path":                       "$bookings",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Add date range filter
	pipeline = append(pipeline, bson.M{
		"$match": bson.M{
			"$or": []bson.M{
				{"bookings": bson.M{"$exists": false}},
				{"bookings.canceled": true},
				{"bookings.fromDate": bson.M{"$gte": *tillDate}},
				{"bookings.tillDate": bson.M{"$lte": *fromDate}},
			},
		},
	})

	// Add room filters
	if len(roomFilters) > 0 {
		pipeline = append(pipeline, bson.M{"$match": roomFilters})
	}

	pipeline = append(pipeline, bson.M{
		"$group": bson.M{
			"_id":         "$name",
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

	// Add pagination
	pipeline = append(pipeline, bson.M{"$skip": (page - 1) * limit})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	// Project to exclude _id and include hotel name
	pipeline = append(pipeline, bson.M{
		"$project": bson.M{
			"_id":            0,
			"name":           1,
			"description":    1,
			"price":          1,
			"capacity":       1,
			"amenities":      1,
			"images":         1,
			"bedType":        1,
			"bedrooms":       1,
			"availableCount": 1,
			"hotelName":      "$hotel.name",
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
