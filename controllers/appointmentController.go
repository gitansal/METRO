package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/gitansal/metro/database"
	models "github.com/gitansal/metro/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/gridfs"
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

		//count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		//defer cancel()
		//if err != nil {
		//	log.Panic(err)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the email"})
		//}
		//if count > 0 {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone already exists"})
		//}

		//password := HashPassword(*user.Password)
		//user.Password = &password

		count, err := appointmentCollection.CountDocuments(ctx, bson.M{"phone": appointment.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "appointment already scheduled for this number"})
		}
		appointment.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		appointment.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		appointment.ID = primitive.NewObjectID()
		//user.User_id = user.ID.Hex()
		//token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		//user.Token = &token
		//user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := appointmentCollection.InsertOne(ctx, appointment)
		if insertErr != nil {
			msg := fmt.Sprintf("appointment was not scheduled")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}
