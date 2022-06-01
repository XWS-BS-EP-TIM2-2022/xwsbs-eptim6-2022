package store

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/consts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	permission := UserPermission{Role: string(consts.ADMIN), Permissions: []string{string(consts.DELETE_USER), string(consts.DELETE_POSTS), string(consts.VIEW_POSTS), string(consts.VIEW_USER), string(consts.VIEW_JOB_OFFER)}}
	permissionUser := UserPermission{Role: string(consts.USER), Permissions: []string{string(consts.CREATE_POSTS), string(consts.VIEW_POSTS), string(consts.UPDATE_POSTS), string(consts.VIEW_USER), string(consts.UPDATE_USER), string(consts.VIEW_JOB_OFFER)}}
	permisionsCompanyOwner := UserPermission{Role: string(consts.COMPANY_OWNER), Permissions: []string{string(consts.CREATE_JOB_OFFER), string(consts.UPDATE_JOB_OFFER), string(consts.DELETE_JOB_OFFER)}}
	ps.AddNew(permission)
	ps.AddNew(permissionUser)
	ps.AddNew(permisionsCompanyOwner)
}
