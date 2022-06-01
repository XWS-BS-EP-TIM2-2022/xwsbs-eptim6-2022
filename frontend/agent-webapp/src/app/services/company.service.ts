import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

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
}
