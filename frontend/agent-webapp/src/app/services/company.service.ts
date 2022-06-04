import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {
  
  constructor(private http: HttpClient) { }

  registerNewCompany(company : any)
  {
    var body = {
      'name' : company.name,
      'address' : company.address,
      'contactInfo' : {
        'email' : company.email,
        'phone' : company.phone,
      },
      'description' : company.description,
      'culture' : company.culture,
      'web' : "",
      'yearOfEstablishment' : company.yearOfEstablishment
    }
    return this.http.post("/api/companies", body, { responseType: 'text' });
  }

  getAll() : Observable<any> {
    return this.http.get<any>('api/companies');
  }

  getByOwner(id : any) : Observable<any> {
    return this.http.get<any>('api/users/' + encodeURIComponent(id) + '/companies');
  }

  getById(id : any) : Observable<any> {
    return this.http.get<any>('api/companies/' + encodeURIComponent(id));
  }

  approveRequest(company : any){
    return this.http.put('/api/companies/' + encodeURIComponent(company.id), company);
  }

  getComments(id : any) : Observable<any[]> {
    return this.http.get<any[]>('api/companies/' + encodeURIComponent(id) + '/comments');
  }

  getSalaries(id : any) : Observable<any[]> {
    return this.http.get<any[]>('api/companies/' + encodeURIComponent(id) + '/salaries');
  }

  getInterviews(id : any) : Observable<any[]> {
    return this.http.get<any[]>('api/companies/' + encodeURIComponent(id) + '/interviews');
  }

  addComment(comment : any, id : any)
  {
    var body = {
      'companyId' : id,
      'creatorId' : localStorage.getItem('id'),
      'experienceLevel' : comment.experienceLevel,
      'positionId' : 1,
      'title' : '',
      'text' : comment.text,
      'rate' : comment.rate
    }
    return this.http.post('api/companies/' + encodeURIComponent(id) + '/comments', body, { responseType: 'text' });
  }

  addSalary(salary : any, id : any)
  {
    var body = {
      'companyId' : id,
      'experienceLevel' : salary.experienceLevel,
      'positionId' : 1,
      'creatorId' : localStorage.getItem('id'),
      'salary' : salary.salary,
    }
    return this.http.post('api/companies/' + encodeURIComponent(id) + '/salaries', body, { responseType: 'text' });
  }

  addInterview(interview : any, id : any)
  {
    var body = {
      'companyId' : id,
      'creatorId' : localStorage.getItem('id'),
      'positionId' : 1,
      'experienceLevel' : interview.experienceLevel,
      'difficultyLevel' : interview.difficultyLevel,
      'selectionProcessDuration' : interview.selectionProcessDuration
    }
    return this.http.post('api/companies/' + encodeURIComponent(id) + '/interviews', body, { responseType: 'text' });
  }

}
