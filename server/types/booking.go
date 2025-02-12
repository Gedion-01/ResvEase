package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	Pending   Status = "pending"
	Confirmed Status = "confirmed"
	Canceled  Status = "canceled"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	FirstName  string             `bson:"firstName" json:"firstName"`
	LastName   string             `bson:"lastName" json:"lastName"`
	Email      string             `bson:"email" json:"email"`
	Phone      string             `bson:"phone" json:"phone"`
	SpecialReq string             `bson:"specialReq" json:"specialReq"`
	NumPersons int                `bson:"numPersons,omitempty" json:"numPersons,omitempty"`
	FromDate   time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate   time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
	Canceled   bool               `bson:"canceled" json:"canceled"`
	Status     Status             `bson:"status" json:"status"`
	Hotel      Hotel              `bson:"hotel" json:"hotel"`
	Room       Room               `bson:"room" json:"room"`
	TotalPrice float64            `bson:"totalPrice" json:"totalPrice"`
	CreatedAt  time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt  time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
