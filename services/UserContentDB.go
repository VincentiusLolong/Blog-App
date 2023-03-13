package services

import (
	"context"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *services) CreateContent(addcontents models.AllContents, str string, a context.Context) (*mongo.InsertOneResult, error) {
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

	result, err := c.monggose.ContentCollection().InsertOne(a, newContent)
	return result, err

}

func (c *services) FindContent(a context.Context, str string) ([]models.AllContents, error) {
	var allcontent []models.AllContents

	objId, _ := primitive.ObjectIDFromHex(str)
	cursor, err := c.monggose.ContentCollection().Find(a, bson.M{"user_id": objId})
	if err != nil {
		return nil, err
	}

	err = cursor.All(a, &allcontent)
	if err != nil {
		return nil, err
	}

	return allcontent, nil
}

func (c *services) ContentEdit(a context.Context, str string, cid string, bsondata primitive.M) (*mongo.UpdateResult, error) {
	objId, _ := primitive.ObjectIDFromHex(str)
	ctjId, _ := primitive.ObjectIDFromHex(cid)
	filter := bson.M{"content_id": ctjId, "user_id": objId}
	update := bson.D{{Key: "$set", Value: bsondata}}

	result, err := c.monggose.ContentCollection().UpdateOne(a, filter, update)
	return result, err

}

func (c *services) ContentDelete(a context.Context, str, content_id string) (*mongo.DeleteResult, error) {
	objId, _ := primitive.ObjectIDFromHex(str)
	ctjId, _ := primitive.ObjectIDFromHex(content_id)
	filter := bson.M{"content_id": ctjId, "user_id": objId}

	result, err := c.monggose.ContentCollection().DeleteOne(a, filter)
	return result, err
}

func LookUpContent()
