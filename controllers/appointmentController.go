package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/gitansal/METRO/database"
	models "github.com/gitansal/METRO/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var appointmentCollection *mongo.Collection = database.OpenCollection(database.Client, "appointment")
var validate = validator.New()

func Schedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var appointment models.Appointment
		defer cancel()
		if err := c.BindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validatorErr := validate.Struct(appointment)
		if validatorErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorErr.Error()})
			return
		}

		count, err := appointmentCollection.CountDocuments(ctx, bson.M{"phone": appointment.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "appointment already scheduled for this number"})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "appointment was not scheduled"})
			return
		}
		appointment.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		appointment.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		appointment.ID = primitive.NewObjectID()
		resultInsertionNumber, insertErr := appointmentCollection.InsertOne(ctx, appointment)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "appointment was not scheduled"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}
