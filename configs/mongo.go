package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB interface {
	UserCollection() *mongo.Collection
	ContentCollection() *mongo.Collection
	CommentsCollection() *mongo.Collection
}

type monggose struct {
	collection *mongo.Collection
}

func New() MongoDB {
	return &monggose{}
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(AllEnv("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = ConnectDB()

// getting database collections
func (m *monggose) UserCollection() *mongo.Collection {
	m.collection = Client.Database(AllEnv("DATABASE")).Collection(AllEnv("THECOLLECTION"))
	return m.collection
}

func (m *monggose) ContentCollection() *mongo.Collection {
	m.collection = Client.Database(AllEnv("DATABASE")).Collection(AllEnv("PXCOLLECTIONS"))
	return m.collection
}

func (m *monggose) CommentsCollection() *mongo.Collection {
	m.collection = Client.Database(AllEnv("DATABASE")).Collection(AllEnv("COMCOLLECTIONS"))
	return m.collection
}
