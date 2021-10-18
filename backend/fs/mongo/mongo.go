package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct{}

var Collection *mongo.Collection
var Context context.Context

func (s Mongo) Write(name string, content []byte) error {
	filename := name + ".txt"
	filter := bson.M{"name": filename}
	update := bson.M{"$set": bson.M{"data": string(content)}}
	opts := options.Update().SetUpsert(true)

	_, err := Collection.UpdateOne(Context, filter, update, opts)
	return err
}

func (s Mongo) Remove(name string) error {
	filename := name + ".txt"
	filter := bson.D{{"name", filename}}
	_, err := Collection.DeleteOne(Context, filter)
	return err
}

func (s Mongo) Read(name string) ([]byte, error) {
	filename := name + ".txt"
	query := bson.D{{"name", filename}}
	result := struct {
		Data string `bson:"data"`
	}{}
	err := Collection.FindOne(Context, query).Decode(&result)
	if err != nil {
		return nil, err
	}
	return []byte(result.Data), nil
}

func (s Mongo) GetAllFileNames() []string {
	var dbs []struct {
		Name string `bson:"name"`
	}
	cursor, _ := Collection.Find(Context, bson.D{})
	fileNames := []string{}

	cursor.All(Context, &dbs)
	for _, db := range dbs {
		fileNames = append(fileNames, db.Name)
	}
	return fileNames
}

func (s Mongo) IsFileExist(name string) bool {
	_, err := Mongo{}.Read(name)
	return err == nil
}
