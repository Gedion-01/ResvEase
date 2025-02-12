package db

import (
	"context"
	"hotel-reservation/types"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(context.Context, bson.M, int, int) ([]*types.Booking, error)
	GetBooking(context.Context, bson.M) (*types.Booking, error)
	GetBookingByID(context.Context, string) (*types.Booking, error)
	UpdateBooking(context.Context, string, bson.M) error
	IsBooked(context.Context, string, string, string) (bool, error)
	Exists(context.Context, bson.M) (bool, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	dbname := os.Getenv(MongoDbNameEnvName)
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(dbname).Collection("bookings"),
	}
}

func (s *MongoBookingStore) UpdateBooking(ctx context.Context, id string, update bson.M) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	m := bson.M{"$set": update}
	_, err = s.coll.UpdateByID(ctx, oid, m)
	return err
}

func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M, page int, limit int) ([]*types.Booking, error) {
	opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	curr, err := s.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var bookings []*types.Booking
	if err := curr.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *MongoBookingStore) GetBooking(ctx context.Context, filter bson.M) (*types.Booking, error) {
	result := s.coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var booking types.Booking
	if err := result.Decode(&booking); err != nil {
		return nil, err
	}
	return &booking, nil
}

func (s *MongoBookingStore) GetBookingByID(ctx context.Context, id string) (*types.Booking, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking types.Booking
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking); err != nil {
		return nil, err
	}
	return &booking, err
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	resp, err := s.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}

func (s *MongoBookingStore) Exists(ctx context.Context, filter bson.M) (bool, error) {
	var result bson.M
	err := s.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *MongoBookingStore) IsBooked(ctx context.Context, roomID, fromDate, tillDate string) (bool, error) {
	where := bson.M{
		"roomID": roomID,
		"fromDate": bson.M{
			"$gte": fromDate,
		},
		"tillDate": bson.M{
			"$lte": tillDate,
		},
		"cancelled": false,
		"status":    types.Confirmed,
	}
	return s.Exists(ctx, where)
}
