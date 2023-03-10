package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string            `json:"first_name" validate:"required"`
	Last_name  *string            `json:"last_name" validate:"required"`
	Phone      *string            `json:"phone" validate:"required,min=10,max=10"`
	Address    *string            `json:"address" validate:"required"`
	Purpose    *string            `json:"purpose" validate:"required"`
	User_id    *string            `json:"user_id" validate:"required"`
	Date       *string            `json:"date" validate:"required"`
	Time       *string            `json:"time" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
