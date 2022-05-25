import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpStatusCode } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { 
    http.options("",{})

  }

  login(user : any) {
    const body = {
      'username': user.username,
      'password': user.password
    };
    
    return this.http.put("/api/auth/session", JSON.stringify(body)).pipe(
      map((res: any) => {
        console.log('Login success');
        localStorage.setItem('jwt', res.token);
        console.log(res);
      })
    );;
  }

  public logOut(): void {
    localStorage.removeItem('jwt');
  }

  register(user : any) {
    const body = {
      'username': user.username,
      'password': user.password,
      'name' : user.name,
      'surname' : user.surname,
      'email' : user.email,
      'role' : "user"
    };
    return this.http.post("/api/auth/users", JSON.stringify(body));
  }

  
  changePassword(req : any) {
    const body = {
      'username': 'petra',
      'oldPassword': req.oldPassword,
      'NewPassword' : req.NewPassword,
    };
    return this.http.put("/api/auth/users/password", JSON.stringify(body));
  }

  getUser(): Observable<any> {
    return this.http.get<Observable<any>>('/api/whoami');
  }
}
