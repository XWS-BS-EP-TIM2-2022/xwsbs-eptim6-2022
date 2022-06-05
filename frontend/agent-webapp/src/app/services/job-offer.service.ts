import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class JobOfferService {

  constructor(private http: HttpClient) { }

  createJobOffer(offer : any, id : any)
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

    return this.http.post("/api/companies/" + encodeURIComponent(id)+ "/job-offers", body, { responseType: 'text' });
  
  }

  getJobOffers(): Observable<any> {
    return this.http.get<Observable<any>>('/api/job-offers');
  }
}
