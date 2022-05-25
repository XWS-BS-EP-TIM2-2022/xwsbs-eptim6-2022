package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Permission string

const (
	CREATE_POSTS Permission = "CREATE_POSTS"
	VIEW_POSTS   Permission = "VIEW_POSTS"
	UPDATE_POSTS Permission = "UPDATE_POSTS"
	DELETE_POSTS Permission = "DELETE_POSTS"

	CREATE_USER Permission = "CREATE_USER"
	UPDATE_USER Permission = "UPDATE_USER"
	VIEW_USER   Permission = "VIEW_USER"
	DELETE_USER Permission = "DELETE_USER"
)

type UserPermission struct {
	Role        string
	Permissions []string
}

type UserPermissions struct {
	Permissions []UserPermission
}

type PermissionsStore struct {
	PermissionsCollection *mongo.Collection
}

func InitPermissionsStore(mongoUri string) *PermissionsStore {
	mongoDbName := "permissions_database"
	mongoCollectionName := "permissions"
	client := CreateMongoDBConnection(mongoUri)
	collection := client.Database(mongoDbName).Collection(mongoCollectionName)
	fmt.Println(collection.Name())
	permissionStore := PermissionsStore{PermissionsCollection: collection}
	//permissionStore.InsertData()
	return &permissionStore
}

//TODO: Onda u gejtveju treba da dodam vezu izmedju permisija i zahtijeva.

func (ps *PermissionsStore) FindByUserRole(role string) (*UserPermission, error) {
	filter := bson.D{{"role", role}}
	var userPermission UserPermission
	err := ps.PermissionsCollection.FindOne(context.TODO(), filter).Decode(&userPermission)
	//cur, err := ps.PermissionsCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	return &userPermission, err
}

func (ps *PermissionsStore) AddNew(permission UserPermission) {
	insertResult, err := ps.PermissionsCollection.InsertOne(context.TODO(), permission)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (ps *PermissionsStore) InsertData() {
	permission := UserPermission{Role: "ADMIN", Permissions: []string{string(DELETE_USER), string(DELETE_POSTS), string(VIEW_POSTS), string(VIEW_USER)}}
	permissionUser := UserPermission{Role: "USER", Permissions: []string{string(CREATE_POSTS), string(VIEW_POSTS), string(UPDATE_POSTS), string(VIEW_USER), string(UPDATE_USER)}}
	ps.AddNew(permission)
	ps.AddNew(permissionUser)
}
