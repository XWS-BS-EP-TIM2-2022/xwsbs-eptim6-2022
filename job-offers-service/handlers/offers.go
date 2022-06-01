package handlers

import (
	"errors"
	"job-offers-service/startup/config"
	"job-offers-service/store"
)

type JobOffersHandler struct {
	jobOffersStore *store.JobOffersStore
}

func NewOffersHandler(config config.Config) *JobOffersHandler {
	return &JobOffersHandler{jobOffersStore: store.InitJobOffersStore(config.MongoDbUri)}
}

func (joh *JobOffersHandler) GetAll() (*[]store.Offer, error) {
	return joh.jobOffersStore.GetAll()
}

func (joh *JobOffersHandler) SearchByPosition(position string) (store.Offers, error) {
	return joh.jobOffersStore.GetByPosition(position)
}

func (joh *JobOffersHandler) CreateJobOffer(in *store.Offer) (store.Offer, error) {
	//currentTime := time.Now().Format("02.01.2006 15:04")
	//offer := store.Offer{CreatedOn: currentTime}
	err := joh.jobOffersStore.CreateJobOffer(*in)
	if err != nil {
		return store.Offer{}, errors.New("Could not create an offer")
	}
	return *in, nil
}
