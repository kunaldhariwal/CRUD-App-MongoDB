package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		BlockId   string        `json:"BlockId" bson:"BlockId"`
		BlockHash string     `json:"BlockHash" bson:"BlockHash"`
		Height    int        `json:"Height" bson:"Height"`
		ActivityId int       `json:"ActivityId" bson:"ActivityId"`
		EventId int          `json:"EventId" bson:"EventId"`
	}
)
