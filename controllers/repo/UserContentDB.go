package repo

import (
	"context"
	"errors"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var contentCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("PXCOLLECTIONS"))

func CreateContent(addcontents models.AllContents, str string, a context.Context) (*mongo.InsertOneResult, error) {
	objId, _ := primitive.ObjectIDFromHex(str)
	newContent := &models.AllContents{
		Content_Id: primitive.NewObjectID(),
		User_id:    objId,
		Title:      addcontents.Title,
		Tags:       addcontents.Tags,
		Type:       addcontents.Type,
		Photos:     addcontents.Photos,
		Content:    addcontents.Content,
	}

	result, err := contentCollection.InsertOne(a, newContent)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func FindContent(a context.Context, str string) ([]models.AllContents, error) {
	var allcontent []models.AllContents

	objId, _ := primitive.ObjectIDFromHex(str)
	cursor, err := contentCollection.Find(a, bson.M{"user_id": objId})
	if err != nil {
		return nil, err
	}

	err = cursor.All(a, &allcontent)
	if err != nil {
		return nil, err
	}

	return allcontent, nil
}

func ContentEdit(a context.Context, str string, cid string, bsondata primitive.M) (*mongo.UpdateResult, error) {
	objId, _ := primitive.ObjectIDFromHex(str)
	ctjId, _ := primitive.ObjectIDFromHex(cid)
	filter := bson.M{"content_id": ctjId, "user_id": objId}
	update := bson.D{{Key: "$set", Value: bsondata}}

	result, err := contentCollection.UpdateOne(a, filter, update)
	if err != nil {
		return nil, errors.New("cant edit")
	} else {
		return result, nil
	}
}
