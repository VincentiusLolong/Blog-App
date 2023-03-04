package repo

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var commentCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("COMCOLLECTIONS"))

func CreateComment(a context.Context, ContentDB models.Comments, userid string) (*mongo.InsertOneResult, error) {
	usersId, _ := primitive.ObjectIDFromHex(userid)
	newComment := &models.Comments{
		Comments_id: primitive.NewObjectID(),
		Content_Id:  ContentDB.Content_Id,
		User_id:     usersId,
		Comment:     ContentDB.Comment,
	}

	result, err := commentCollection.InsertOne(a, newComment)
	return result, err
}

func EditComment(a context.Context, commentid string, userid string, bsondata primitive.M) (*mongo.UpdateResult, error) {
	commId, _ := primitive.ObjectIDFromHex(commentid)
	uId, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.M{"content_id": commId, "user_id": uId}
	update := bson.D{{Key: "$set", Value: bsondata}}

	result, err := contentCollection.UpdateOne(a, filter, update)
	return result, err
}

func CommentHistory() {}

// get by rank :
// 1. User
// 2. Comment_id
// 3.
