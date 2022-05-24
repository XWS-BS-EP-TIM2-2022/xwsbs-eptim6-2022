import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  login(user: any) {
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

  register(user: any) {
    const body = {
      'username': user.username,
      'password': user.password,
      'name': user.name,
      'surname': user.surname,
      'email': user.email,
      'role': "user"
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
  forgotPassword(email: string) {
    const body = { 'email': email };
    return this.http.post("/api/auth/users/forgot-password", JSON.stringify(body), { responseType: 'text' });
  }

  setPassword(password: string, token: string) {
    const body = {
      'password': password,
      'verificationToken': token
    };

    return this.http.post('/api/auth/users/reset-password', JSON.stringify(body), { responseType: 'text' });
  }
}
