import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class JobOffersService {

  constructor(private http: HttpClient) { }

  getJobOffers(): Observable<any> {
    return this.http.get<Observable<any>>('/api/job-offers');
  }
}
