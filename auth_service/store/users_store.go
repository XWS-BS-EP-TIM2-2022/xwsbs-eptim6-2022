package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type UsersStore struct {
	UsersCollection *mongo.Collection
}

type AuthStore interface {
	AddNew(user *User)
	FindByUsername(username string) (User, error)
	FindAll() []User
}

func (us *UsersStore) AddNew(u *User) {
	insertResult, err := us.UsersCollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (us *UsersStore) FindByUsername(username string) (User, error) {
	var user User
	filter := bson.D{{"username", username}}
	err := us.UsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}
func (us *UsersStore) FindAll() []User {
	fmt.Println("FindAll users_store")
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

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
	//mongoUri := "localhost:27017" //os.Getenv("MONGODB_URI")
	//clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//return client, err
}

func InitUsersStore(client *mongo.Client) *UsersStore {
	//mongoUri := "localhost:27017" //os.Getenv("MONGODB_URI")
	//clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//
	//// Check the connection
	//err = client.Ping(context.TODO(), nil)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("users_database").Collection("users")
	fmt.Println(collection.Name())
	return &UsersStore{UsersCollection: collection}
}
