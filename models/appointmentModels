package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Appointment struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string            `json:"full_name" validate:"required"`
	Last_name  *string            `json:"last_name" validate:"required"`
	Phone      *string            `json:"phone" validate:"required"`
	Address    *string            `json:"address" validate:"required"`
	Purpose    *string            `json:"purpose" validate:"required"`
	//Appoint_date     primitive.DateTime            `json:"appoint_date" validate:"required"`
	//Appoint_time     *string            `json:"appoint_time" validate:"required"`
	//Refresh_token  *string   `json:"refresh_token"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	//Appointment_id string    `json:"appointment_id"`
}
