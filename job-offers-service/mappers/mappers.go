package mappers

import (
	jobOffersPb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/job_offers_service"
	"job-offers-service/store"
	"strconv"
)

func stringToInt(s string) int {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return atoi
}
func MapToPb(offer store.Offer) *jobOffersPb.JobOffer {
	return &jobOffersPb.JobOffer{Username: offer.User,
		Id: offer.ID.Hex(), CreatedOn: offer.CreatedOn, Description: offer.Description, Position: offer.Position,
		ValidUntil: offer.ValidUntil, WorkScheduleHours: strconv.Itoa(offer.WorkSchedule.HoursPerWeek), WorkScheduleTitle: offer.WorkSchedule.Title, JobOfferUrl: offer.JobOfferUrl, Experience: offer.Preconditions.Experience}
}
func MapToStore(offer *jobOffersPb.CreateJobOffer) store.Offer {
	return store.Offer{
		CreatedOn: offer.CreatedOn, Description: offer.Description, Position: offer.Position, ValidUntil: offer.ValidUntil,
		WorkSchedule: store.WorkSchedule{Title: offer.WorkScheduleTitle, HoursPerWeek: stringToInt(offer.WorkScheduleHours)}, JobOfferUrl: offer.JobOfferUrl,
		Preconditions: store.Precondition{Experience: offer.Experience}}
}
func MapToResponses(offers *[]store.Offer) *jobOffersPb.JobOffersResponse {
	var response []*jobOffersPb.JobOffer
	for _, offer := range *offers {
		mapped := MapToPb(offer)
		response = append(response, mapped)
	}
	return &jobOffersPb.JobOffersResponse{Offers: response}
}
func MapToResponse(offer store.Offer) *jobOffersPb.JobOfferResponse {
	return &jobOffersPb.JobOfferResponse{Offer: MapToPb(offer)}
}
