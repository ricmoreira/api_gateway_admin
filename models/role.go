package models

import "gopkg.in/mgo.v2/bson"

type Role struct {
	ID                 bson.ObjectId  `json:"id,omitempty" bson:"_id"`
	Role string   `json:"role" valid:"required~Field role cannot be empty" bson:"role"`
	Level    int   `json:"level" valid:"required~Field level cannot be empty or is missing" bson:"level"`
}
