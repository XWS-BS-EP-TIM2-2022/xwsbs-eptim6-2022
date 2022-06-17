package handlers

import (
	"errors"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	"job-offers-service/startup/config"
	"job-offers-service/store"
	"reflect"
	"runtime"
)

type JobOffersHandler struct {
	jobOffersStore *store.JobOffersStore
	log            *logger.LoggerWrapper
}

func NewOffersHandler(config config.Config, wrapper *logger.LoggerWrapper) *JobOffersHandler {
	return &JobOffersHandler{jobOffersStore: store.InitJobOffersStore(config.MongoDbUri), log: wrapper}
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

func getComponentName(methode interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(methode).Pointer()).Name()
}
