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
}
