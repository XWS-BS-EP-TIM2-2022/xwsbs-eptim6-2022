package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Name     string             `json:"name" bson:"name"`
	Surname  string             `json:"surname" bson:"surname"`
	Password string             `json:"password" bson:"password"`
}

type UsersStore struct {
	UsersCollection *mongo.Collection
}

func (us *UsersStore) FindOne(id primitive.ObjectID) (User, error) {
	var user User
	filter := bson.D{{Key: "_id", Value: id}}
	err := us.UsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (us *UsersStore) FindAll() []User {
	cur, err := us.UsersCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, elem)
	}
	cur.Close(context.TODO())
	fmt.Println(users)
	return users
}

func (us *UsersStore) AddUser(u User) *mongo.InsertOneResult {
	insertResult, err := us.UsersCollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return insertResult
}

func InitUsersStore() *UsersStore {
	mongoUri := "localhost:27017" //os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("users_database").Collection("users")
	fmt.Println(collection.Name())
	return &UsersStore{UsersCollection: collection}
}
