package db

import (
	"context"
	"hotel-reservation/types"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]map[string]interface{}, error)
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

// func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
// 	resp, err := s.coll.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var rooms []*types.Room
// 	if err := resp.All(ctx, &rooms); err != nil {
// 		return nil, err
// 	}
// 	return rooms, nil

// }

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]map[string]interface{}, error) {
	pipeline := []bson.M{
		{"$match": filter},
		{"$group": bson.M{
			"_id": "$name",
			"rooms": bson.M{
				"$push": "$$ROOT",
			},
			"availableCount": bson.M{
				"$sum": bson.M{
					"$cond": []interface{}{"$available", 1, 0},
				},
			},
		}},
	}

	cursor, err := s.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []struct {
		ID             string        `bson:"_id"`
		Rooms          []*types.Room `bson:"rooms"`
		AvailableCount int           `bson:"availableCount"`
	}

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var aggregatedResults []map[string]interface{}
	for _, result := range results {
		if len(result.Rooms) > 0 {
			room := result.Rooms[0]
			aggregatedResults = append(aggregatedResults, map[string]interface{}{
				"name":           room.Name,
				"description":    room.Description,
				"price":          room.Price,
				"capacity":       room.Capacity,
				"amenities":      room.Amenities,
				"images":         room.Images,
				"bedType":        room.BedType,
				"bedrooms":       room.Bedrooms,
				"availableCount": result.AvailableCount,
				"hotelID":        room.HotelID,
			})
		}
	}

	return aggregatedResults, nil
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
