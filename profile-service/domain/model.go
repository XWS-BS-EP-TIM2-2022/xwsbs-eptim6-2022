package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppUser struct {
	Id      primitive.ObjectID
	Name    string
	Surname string
}
