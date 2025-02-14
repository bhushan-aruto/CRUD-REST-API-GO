package repository

import (
	"context"
	"errors"
	"time"

	"github.com/bhushan-aruto/ice_cream_shop_rest_api/config"
	"github.com/bhushan-aruto/ice_cream_shop_rest_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var flavourCollection = config.DB.Collection("flavours")

var flavourCollection *mongo.Collection

func InitFlavourRepo() {
	if config.DB == nil {
		panic("Database is not initialized. Call config.ConnectMongo() first!")
	}
	flavourCollection = config.DB.Collection("flavours")
}

func GetAllFlavours() ([]models.Flavour, error) {
	var flavours []models.Flavour
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()

	cursor, err := flavourCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var flavour models.Flavour
		if err := cursor.Decode(&flavour); err != nil {
			return nil, err
		}

		flavours = append(flavours, flavour)
	}
	return flavours, nil
}

func AddFlavour(flavour models.Flavour) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	flavour.ID = primitive.NewObjectID()
	_, err := flavourCollection.InsertOne(ctx, flavour)
	return err

}

func GetFlavourById(id string) (models.Flavour, error) {
	var flavour models.Flavour
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Flavour{}, errors.New("invalid id format")
	}
	err = flavourCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&flavour)
	if err != nil {
		return models.Flavour{}, errors.New("flavor not found")
	}
	return flavour, nil
}

func Updateflavour(id string, updateFlavour models.Flavour) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	result, err := flavourCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updateFlavour})

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("flavours not found or no chanage made")
	}
	return nil
}

func DeleteFlavour(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")

	}
	result, err := flavourCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("flavour not found")
	}
	return nil
}
