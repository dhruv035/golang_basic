package controllers

import (
	"basic_server/configs"
	"basic_server/models"
	"basic_server/responses"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "Gas Tracker")

func GetData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//validate the request body
	filter := bson.D{{}}

	var result2 []models.GasObject
	result, err := userCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Msg: "Failed to fetch data"})
	}

	if err2 := result.All(context.TODO(), &result2); err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Msg: "Failed to transform data"})
	}

	return c.Status(http.StatusCreated).JSON(responses.GasResponse{GasObjects: result2})
}
