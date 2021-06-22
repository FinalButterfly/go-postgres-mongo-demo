package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName = "gonews"
)

// Хранилище данных.
type Store struct {
	db *mongo.Client
}

// Конструктор объекта хранилища.
func New(connectionString string) (*Store, error) {
	mongoOpts := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: client,
	}
	return &s, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection("posts")
	filter := bson.M{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var p storage.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		data = append(data, p)
	}
	return data, cur.Err()
}

func (s *Store) AddPost(post storage.Post) (interface{}, error) {
	collection := s.db.Database(databaseName).Collection("posts")
	cur, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return 0, err
	}
	return cur.InsertedID, err
}

func (s *Store) UpdatePost(post storage.Post) error {
	collection := s.db.Database(databaseName).Collection("posts")
	_, err := collection.UpdateByID(context.Background(), post.Id, bson.M{"$set": post})
	if err != nil {
		return err
	}
	return err
}

func (s *Store) DeletePost(id interface{}) error {
	collection := s.db.Database(databaseName).Collection("posts")
	p, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", id))
	if err != nil {

		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": p})
	if err != nil {
		return err
	}
	return err
}
