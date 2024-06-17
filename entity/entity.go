package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDGenerates struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UpdateAt time.Time          `bson:"updated_at"`
	MQLID    string             `bson:"mqlid"`
	VPSName  string             `bson:"vps_name"`
}
