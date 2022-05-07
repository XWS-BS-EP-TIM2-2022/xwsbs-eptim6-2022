package store

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type UsersStore struct {
	UsersCollection *mongo.Collection
	//AuthHandler     *handlers.AuthHandler
}

type AuthStore interface {
	AddNew(user User)
	FindByUsername(username string) (User, error)
	LoginUser(user User) (string, error)
	FindAll() []User
}

/*func (us *UsersStore) AddNewUser(user User) {
	us.AuthHandler.AddNewUser(user)
}*/

func (us *UsersStore) LoginUser(user User) (string, error) {
	fmt.Println("LoginUser users_store")
	dbUser, err := us.FindByUsername(user.Username)
	if err != nil {
		fmt.Println("User not found")
		//w.WriteHeader(http.StatusNotFound)
		return "", err
	}
	if dbUser.Password == user.Password {
		fmt.Println("generisanje tokena:")
		tokenStr, err := GenerateJWT(dbUser)
		fmt.Println("Token: " + tokenStr)
		//tokenStr := "3424234323"
		if err != nil {
			fmt.Printf("Token generation failed %s\n", err.Error())
			return "", err
		}
		return tokenStr, nil
		//json.NewEncoder(w).Encode(JWT{Token: tokenStr})
	} else {
		fmt.Println("Login failed")
		//w.WriteHeader(http.StatusBadRequest)
		return "", nil
	}
}

func (us *UsersStore) AddNew(u User) {

	if _, err := us.FindByUsername(u.Username); err == nil {
		//w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(ErrorMessage{Message: "Username already in use"})
		fmt.Println("Username already exists")
		return
	}
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

func GenerateJWT(dbUser User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = dbUser.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenStr, err := token.SignedString(secretString)
	if err != nil {
		fmt.Errorf("token signing error")
		return "", err
	}
	return tokenStr, nil
}

var secretString = []byte("secret_key") //TODO: Use ENV Variable
