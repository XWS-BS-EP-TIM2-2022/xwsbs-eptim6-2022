import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class JobOfferService {

  constructor(public _http: HttpClient) { }

  getJobOffers(): Observable<any> {
    return this._http.get<Observable<any>>('/api/job-offers');
  }
}
