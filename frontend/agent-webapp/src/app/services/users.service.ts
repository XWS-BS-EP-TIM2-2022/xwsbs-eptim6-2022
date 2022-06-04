import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpStatusCode } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class UsersService {

  constructor(private http: HttpClient) { }

  login(user: any) {
    const body = {
      'username': user.username,
      'password': user.password,
    };
    return this.http.put("/api/users/session", body).pipe(
      map((res: any) => {
        console.log('Login success');
        localStorage.setItem('jwt', res.jwt);
        localStorage.setItem('role', res.role);
        localStorage.setItem('id', res.id);
        console.log(res);
      })
    );;
  }

  public logOut(): void {
    localStorage.removeItem('jwt');
    localStorage.removeItem('role');
    localStorage.removeItem('id');
  }

  public getUser(id : string) : Observable<any> {
      return this.http.get<Observable<any>>('/api/users/' + encodeURIComponent(id));
  }

  signup(user: any) {
    const body = {
      'username': user.username,
      'name': user.name,
      'surname': user.surname,
      'email': user.email,
      'phone': user.phone,
      'password': user.password,
      'role': user.role
    };
    return this.http.post("/api/users", body);
  }
}
