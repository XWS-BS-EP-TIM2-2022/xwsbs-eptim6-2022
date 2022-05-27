package handlers

import (
	"errors"
	"job-offers-service/store"
	"time"
)

type JobOffersHandler struct {
	jobOffersStore *store.JobOffersStore
}

func NewOffersHandler() *JobOffersHandler {
	//TODO config.MongoUri
	jobOffersStore := store.InitPostsStore("")
	return &JobOffersHandler{jobOffersStore}
}

func (joh *JobOffersHandler) GetAll() (store.Offers, error) {
	return joh.jobOffersStore.GetAll()
}

func (joh *JobOffersHandler) SearchByPosition(position string) (store.Offers, error) {
	return joh.jobOffersStore.GetByPosition(position)
}

//TODO
func (joh *JobOffersHandler) CreateJobOffer(in *store.Offer) (store.Offer, error) {
	currentTime := time.Now().Format("02.01.2006 15:04")
	offer := store.Offer{CreatedOn: currentTime}
	err := joh.jobOffersStore.CreateJobOffer(offer)
	if err != nil {
		return store.Offer{}, errors.New("Could not create an offer")
	}
	return offer, nil
}
