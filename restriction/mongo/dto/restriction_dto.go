package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestrictionDTO struct {
	ID    primitive.ObjectID      `bson:"_id,omitempty"`
	Key   string                  `bson:"key"`
	Value string                  `bson:"value"`
	Rules map[string]TimeRangeDTO `bson:"rules"`
}

type TimeRangeDTO struct {
	StartDate time.Time `bson:"start_date"`
	EndDate   time.Time `bson:"end_date"`
}
