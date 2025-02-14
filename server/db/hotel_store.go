package db

import (
	"context"
	"fmt"
	"hotel-reservation/types"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, Map, Map) error
	GetHotels(context.Context, Map, *Pagination, float64, float64) ([]*types.Hotel, error)
	GetHotelByID(context.Context, string) (*types.Hotel, error)
	CountHotels(context.Context, Map) (int64, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	dbname := os.Getenv(MongoDbNameEnvName)
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbname).Collection("hotels"),
	}
}

func (s *MongoHotelStore) GetHotelByID(ctx context.Context, id string) (*types.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var hotel *types.Hotel
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel); err != nil {
		return nil, err
	}
	return hotel, nil
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter Map, page *Pagination, minPrice, maxPrice float64) ([]*types.Hotel, error) {
	pipeline := []bson.M{
		{"$match": filter},
		{"$lookup": bson.M{
			"from":         "rooms",
			"localField":   "_id",
			"foreignField": "hotelID",
			"as":           "rooms",
		}},
		{"$addFields": bson.M{
			"prices": bson.M{
				"$reduce": bson.M{
					"input":        "$rooms.price",
					"initialValue": []interface{}{},
					"in": bson.M{
						"$cond": bson.M{
							"if":   bson.M{"$in": []interface{}{"$$this", "$$value"}},
							"then": "$$value",
							"else": bson.M{"$concatArrays": []interface{}{"$$value", []interface{}{"$$this"}}},
						},
					},
				},
			},
		}},
	}

	if minPrice > 0 || maxPrice > 0 {
		priceRangeFilter := bson.M{}
		if minPrice > 0 {
			priceRangeFilter["$gte"] = minPrice
		}
		if maxPrice > 0 {
			priceRangeFilter["$lte"] = maxPrice
		}
		pipeline = append(pipeline, bson.M{
			"$addFields": bson.M{
				"prices": bson.M{
					"$filter": bson.M{
						"input": "$prices",
						"as":    "price",
						"cond": bson.M{"$and": []bson.M{
							{"$gte": []interface{}{"$$price", minPrice}},
							{"$lte": []interface{}{"$$price", maxPrice}},
						}},
					},
				},
			},
		})
	}

	pipeline = append(pipeline, bson.M{
		"$match": bson.M{
			"prices": bson.M{"$ne": []interface{}{}},
		},
	})

	pipeline = append(pipeline, []bson.M{
		{"$project": bson.M{
			"rooms": 0,
		}},
		{"$skip": (page.Page - 1) * page.Limit},
		{"$limit": page.Limit},
	}...)

	cursor, err := s.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)

	var hotels []*types.Hotel
	if err := cursor.All(ctx, &hotels); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(hotels) == 0 {
		fmt.Println("hotel resource not found")
		return nil, fmt.Errorf("hotel resource not found")
	}

	return hotels, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, filter Map, update Map) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) CountHotels(ctx context.Context, filter Map) (int64, error) {
	return s.coll.CountDocuments(ctx, filter)
}
