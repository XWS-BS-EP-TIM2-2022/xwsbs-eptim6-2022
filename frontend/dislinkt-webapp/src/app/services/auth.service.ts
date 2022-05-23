import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  login(user : any) {
    const body = {
      'username': user.username,
      'password': user.password
    };
    return this.http.put("http://localhost:8081/api/auth/session", JSON.stringify(body));
  }

  register(user : any) {
    const body = {
      'username': user.username,
      'password': user.password,
      'name' : user.name,
      'surname' : user.surname
    };
    return this.http.post("http://localhost:8081/api/auth/users", JSON.stringify(body));
  }
}
