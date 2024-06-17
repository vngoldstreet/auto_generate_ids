package service

import (
	"context"
	"errors"
	"log"
	"time"
	"vietvd/gennate_id/entity"
	"vietvd/gennate_id/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "FXGHDatabases"
	collectionName = "id_auto_generates"
)

func DeleteMongoByID(id string) error {
	client := repository.GetMongoClient()

	collection := client.Database(databaseName).Collection(collectionName)
	idObj, err := primitive.ObjectIDFromHex(id) // Replace with your actual document ID
	if err != nil {
		return err
	}

	// Create a filter to match the document by _id
	filter := bson.M{"_id": idObj}

	// Delete the document
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	log.Printf("Deleted %d document(s)\n", deleteResult.DeletedCount)
	return nil
}

func UpMongoByID(id string, vps_name string) error {
	client := repository.GetMongoClient()
	collection := client.Database(databaseName).Collection(collectionName)

	idObj, err := primitive.ObjectIDFromHex(id) // Replace with your actual document ID
	if err != nil {
		return err
	}

	// Define the update document
	update := bson.M{
		"$set": bson.M{
			"vps_name":   vps_name,
			"updated_at": time.Now(),
		},
	}

	// Update the document
	updateResult, err := collection.UpdateByID(context.Background(), idObj, update)
	if err != nil {
		return err
	}

	log.Printf("Matched %d document(s) and modified %d document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func SetMongoDB(mql_id string, vps_name string) error {
	client := repository.GetMongoClient()
	collection := client.Database(databaseName).Collection(collectionName)
	document := entity.IDGenerates{
		UpdateAt: time.Now(),
		MQLID:    mql_id,
		VPSName:  vps_name,
	}

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	return nil
}

func GetMongoDB(vps_name string) (entity.IDGenerates, error) {
	client := repository.GetMongoClient()
	collection := client.Database(databaseName).Collection(collectionName)
	filter := bson.M{"vps_name": vps_name} // Replace with your actual filter condition

	var result entity.IDGenerates
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.IDGenerates{}, errors.New("no document found")
		} else {
			return entity.IDGenerates{}, err
		}
	}
	return result, nil
}
