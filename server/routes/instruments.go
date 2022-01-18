package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// validates data received from endpoint
var validate = validator.New()

// initiate connection
var instCollection *mongo.Collection = connectCollection(Client, "instruments")

// adds instruments to collection
func AddInstruments(c *gin.Context) {

	// create context with timeout
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var instrument models.Instrument

    if err := c.BindJSON(&instrument); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// if error, return error as JSON
	validationErr := validate.Struct(instrument)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	instrument.ID = primitive.NewObjectID()

	// add instrument to DB
	result, insertErr := instCollection.InsertOne(ctx, instrument)

	if insertErr != nil {
		Errmsg := fmt.Sprintf("Instrument was not added.")
		c.JSON(http.StatusInternalServerError, gin.H{"error": Errmsg})
		fmt.Println(insertErr)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result)
}

// gets instruments from collection	
func GetInstruments(c *gin.Context){
	
	// create context with timeout
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	// create array of instruments
	var instruments []bson.M

	// get data from DB
	cursor, err := instCollection.Find(ctx, bson.M{})
	
	// handle errors
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(ctx, &instruments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	defer cancel()
	
	fmt.Println(instruments)

	c.JSON(http.StatusOK, instruments)
}

// get instrument by ID 
func GetInstrumentsById(c *gin.Context){

	instrumentID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(instrumentID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	var instrument bson.M

    if err := instCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&instrument); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)

        return
    }

    defer cancel()

    fmt.Println(instrument)

	c.JSON(http.StatusOK, instrument)
}

// update existing instrument by ID
func UpdateInstruments(c *gin.Context) {

	instrumentID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(instrumentID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	var instrument models.Instrument

	if err := c.BindJSON(&instrument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// validation
	validationErr := validate.Struct(instrument)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	// update instrument by ID
	result, err := instCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"type":  instrument.Type,
			"name": instrument.Name,
			"price": instrument.Price,
			"quantity": instrument.Quantity,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)
}

// delete instrument by id
func DeleteInstruments(c * gin.Context){
	instrumentID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(instrumentID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// delete instrument by ID
	result, err := instCollection.DeleteOne(ctx, bson.M{"_id": docID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)

		return
    }
	defer cancel()
    
	c.JSON(http.StatusOK, result.DeletedCount)
}