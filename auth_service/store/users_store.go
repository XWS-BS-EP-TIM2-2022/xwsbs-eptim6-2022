package store

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"reflect"
	"time"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username          string    `json:"username" validate:"required,lt=50"`
	Name              string    `json:"name" validate:"required"`
	Surname           string    `json:"surname" validate:"required"`
	Password          string    `json:"password" validate:"required"`
	Email             string    `json:"email" validate:"required,email"`
	Role              string    `json:"role" validate:"required"`
	FailedLogins      int       `json:"failed-logins" bson:"failed-logins"`
	Blocked           bool      `json:"blocked"`
	BlockedUntil      time.Time `json:"blocked-until" bson:"blocked-until"`
	IsActivated       bool      `json:"isActivated" bson:"is-activated"`
	VerificationToken string    `json:"verificationToken" bson:"verification-token"`
	TokenExpiration   time.Time `json:"tokenExpiration" bson:"token-expiration"`
}
type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old-password"`
	NewPassword string `json:"new-password"`
}
type UsersStore struct {
	UsersCollection *mongo.Collection
}

var validate *validator.Validate

func validateUserData(user User) (*User, error) {

	err := validate.Struct(user)
	if err != nil {
		return nil, err
	}
	value := reflect.ValueOf(&user).Elem()

	// loop over the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		// check if the field is a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		// set field to escaped version of the string
		field.SetString(html.EscapeString(str))
	}
	return &user, nil
}

func (us *UsersStore) AddNew(u *User) error {
	user, err := validateUserData(*u)
	if err != nil {
		return err
	}
	insertResult, err := us.UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func (us *UsersStore) FindByUsername(username string) (User, error) {
	var user User
	filter := bson.D{{"username", username}}
	err := us.UsersCollection.FindOne(context.TODO(), filter).Decode(&user)
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

func (us *UsersStore) UpdateFailedLogForUser(username string) error {
	filter := bson.D{{"username", username}}

	update := bson.D{
		{"$inc", bson.D{
			{"failed-logins", 1},
		}},
	}

	_, err := us.UsersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (us *UsersStore) ResetFailedLogForUser(username string) error {
	filter := bson.D{{"username", username}}

	update := bson.D{
		{"$set", bson.D{
			{"failed-logins", 0},
			{"blocked", false},
		}},
	}

	_, err := us.UsersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (us *UsersStore) BlockUser(username string) error {
	filter := bson.D{{"username", username}}

	update := bson.D{
		{"$set", bson.D{
			{"blocked", true},
			{"blocked-until", time.Now().AddDate(0, 0, 1)},
		}},
	}
	_, err := us.UsersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (us *UsersStore) UpdatePassword(username string, password string) error {
	filter := bson.D{{"username", username}}

	update := bson.D{
		{"$set", bson.D{
			{"password", password},
		}},
	}
	_, err := us.UsersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func InitUsersStore(config config.Config) *UsersStore {
	validate = validator.New()
	client := CreateMongoDBConnection(config.MongoDbUri)
	collection := client.Database(config.MongoDbName).Collection(config.MongoDbCollection)
	fmt.Println(collection.Name())
	return &UsersStore{UsersCollection: collection}
}

func (us *UsersStore) FindByToken(token string) User {
	for _, element := range us.FindAll() {
		err := bcrypt.CompareHashAndPassword([]byte(element.VerificationToken), []byte(token))
		if err == nil {
			return element
		}
	}

	return User{}
}

func (us *UsersStore) ActivateAccount(username string) error {
	filter := bson.D{{"username", username}}

	update := bson.D{
		{"$set", bson.D{
			{"is-activated", true},
		}},
	}
	_, err := us.UsersCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (us *UsersStore) FindByEmail(email string) (User, error) {
	var user User
	filter := bson.D{{"email", email}}
	err := us.UsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func (us *UsersStore) RefreshToken(username string, token string) error {
	filter := bson.D{{"username", username}}

	update := bson.D{
		{"$set", bson.D{
			{"verification-token", token},
			{"token-expiration", time.Now().Add(time.Hour * 2)},
		}},
	}
	_, err := us.UsersCollection.UpdateOne(context.TODO(), filter, update)
	return err
}
