import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { Education, Experience, Interest, Skill, UpdateUser } from '../model/profile';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(public _http: HttpClient) { }

  getUser(): Observable<any> {
    return this._http.get<Observable<any>>('/api/whoami');
  }

  updateUser(user: UpdateUser): Observable<any> {
    return this._http.put<Observable<any>>('http://localhost:5000/', user);
  }

  addNewExperience(experience: Experience): Observable<any> {
    return this._http.put<any>('/api/users/experience', experience)
    .pipe(
      tap(data => console.log("experience: ", experience))
    )
  }

  addNewEducation(education: Education): Observable<any> {
    return this._http.put<any>('api/users/education', education)
    .pipe(
      tap(data => console.log("data: ", data))
    )
  }

  addNewSkill(skill: Skill): Observable<any> {
    return this._http.put<any>('api/users/skill', skill)
    .pipe(
      tap(data => console.log("data: ", data))
    )
  }

  addNewInterest(interest: Interest): Observable<any> {
    return this._http.put<any>('api/users/interest', interest)
    .pipe(
      tap(data => console.log("data: ", data))
    )
  }
}
