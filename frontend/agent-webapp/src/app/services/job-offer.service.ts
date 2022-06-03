import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class JobOfferService {

  constructor(private http: HttpClient) { }

  createJobOffer(offer : any)
  {
    var body ={
      'position':{
          'name': offer.position
      },
      'shareOnDislinkt': offer.shareOnDislinkt,
      'validTo': offer.validTo,
      'createdAt':"",
      'description': offer.description ,
      'experience': offer.experience,
      'workSchedule':{
          'name': offer.workSchedule,
          'hoursPerWeek': offer.hoursPerWeek
      }
  }
    return this.http.post("/api/companies/" + encodeURIComponent("1")+ "/job-offers", body, { responseType: 'text' });
  }
}
