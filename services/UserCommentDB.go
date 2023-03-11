package services

import (
	"context"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *services) CreateComment(a context.Context, ContentDB models.Comments, userid string) (*mongo.InsertOneResult, error) {
	usersId, _ := primitive.ObjectIDFromHex(userid)
	newComment := &models.Comments{
		Comments_id: primitive.NewObjectID(),
		Content_Id:  ContentDB.Content_Id,
		User_id:     usersId,
		Comment:     ContentDB.Comment,
	}

	result, err := c.monggose.CommentsCollection().InsertOne(a, newComment)
	return result, err
}

func (c *services) EditComment(a context.Context, commentid string, userid string, bsondata primitive.M) (*mongo.UpdateResult, error) {
	uId, _ := primitive.ObjectIDFromHex(userid)
	mId, _ := primitive.ObjectIDFromHex(commentid)
	filter := bson.M{"comments_id": mId, "user_id": uId}
	update := bson.D{{Key: "$set", Value: bsondata}}

	result, err := c.monggose.CommentsCollection().UpdateOne(a, filter, update)
	return result, err
}

//

// get by rank :
// 1. User
// 2. Comment_id
// 3.
