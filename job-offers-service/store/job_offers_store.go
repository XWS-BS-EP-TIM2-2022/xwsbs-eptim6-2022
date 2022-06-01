package store

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
)

type Offer struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User          string             `bson:"user"`
	Position      string             `bson:"position"`
	Description   string             `bson:"description"`
	Preconditions Precondition       `bson:"preconditions"`
	CreatedOn     string             `bson:"createdOn"`
	ValidUntil    string
	WorkSchedule  WorkSchedule
	JobOfferUrl   string
}
type WorkSchedule struct {
	Title        string
	HoursPerWeek int
}
type Precondition struct {
	Experience string   `bson:"experience"`
	Educations []string `bson:"educations"`
	Skills     []string `bson:"skills"`
}

type JobOffersStore struct {
	JobOffersCollection *mongo.Collection
}

type Offers []*Offer

func (p *Offers) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Offer) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Offer) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (ps *JobOffersStore) GetAll() (*[]Offer, error) {
	cur, err := ps.JobOffersCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	var offers []Offer
	for cur.Next(context.TODO()) {
		var elem Offer
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		offers = append(offers, elem)
	}
	cur.Close(context.TODO())
	return &offers, nil
}

func (ps *JobOffersStore) GetByPosition(position string) (Offers, error) {
	filter := bson.D{{"position", position}}
	cur, err := ps.JobOffersCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return nil, err
	}
	var offers Offers
	for cur.Next(context.TODO()) {
		var elem Offer
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		offers = append(offers, &elem)
	}
	cur.Close(context.TODO())
	return offers, nil
}

func (ps *JobOffersStore) CreateJobOffer(newOffer Offer) error {
	_, err := ps.JobOffersCollection.InsertOne(context.TODO(), newOffer)
	if err != nil {
		return err
	}
	return nil
}

func InitJobOffersStore(mongoUri string) *JobOffersStore {
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("job_offers_db").Collection("job_offers")
	fmt.Println(collection.Name())
	return &JobOffersStore{JobOffersCollection: collection}
}
