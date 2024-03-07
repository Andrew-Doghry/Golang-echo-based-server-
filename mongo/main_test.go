package mongo

import (
	"context"
	"testing"

	"github.com/tj/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users []User = []User{
	{FName: "john", LName: "doghry"},
	{FName: "andrew", LName: "doghry"},
	{FName: "peter", LName: "doghry"},
}

func TestInsertDsata(t *testing.T) {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("error in connection => %v", err)
	}
	db := c.Database("tronics")
	collection := db.Collection("products")
	res, err := insertData(*collection, []interface{}{users})
	assert.Nil(t, err)
	assert.IsType(t, mongo.InsertManyResult{}, res)
	t.Log("insertedIds", res.InsertedIDs)
}
