package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	Data      string    `bson:"data,omitempty" json:"data,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func (l *LogEntry) Insert(entry LogEntry, ctx context.Context) error {
	collection := client.Database("logs").Collection("logs")
	var now = time.Now()
	if _, err := collection.InsertOne(ctx, LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: now,
		UpdatedAt: now,
	}); err != nil {
		return err
	}
	return nil
}

func (l *LogEntry) All(ctx context.Context) ([]LogEntry, error) {
	collection := client.Database("logs").Collection("logs")
	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []LogEntry
	for cursor.Next(ctx) {
		var item LogEntry
		if err := cursor.Decode(&item); err != nil {
			log.Printf("Error decoding the cursor \n%v\n", err)
			return nil, err
		}
		logs = append(logs, item)
	}
	return logs, nil
}

func (l *LogEntry) GetOne(ctx context.Context, id string) (*LogEntry, error) {
	collection := client.Database("logs").Collection("logs")
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entry LogEntry
	if err := collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry); err != nil {
		return nil, err
	}
	return &entry, nil
}

func (l *LogEntry) DropCollection(ctx context.Context) error {
	collection := client.Database("logs").Collection("logs")
	if err := collection.Drop(ctx); err != nil {
		return err
	}
	return nil
}

func (l *LogEntry) Update(ctx context.Context) (*mongo.UpdateResult, error) {
	collection := client.Database("logs").Collection("logs")
	docId, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		return nil, err
	}
	return collection.UpdateOne(ctx, bson.M{"_id": docId}, bson.D{
		bson.E{"$set", bson.D{
			bson.E{"name", l.Name},
			bson.E{"data", l.Data},
			bson.E{"updated_at", time.Now()},
		}},
	})
}
