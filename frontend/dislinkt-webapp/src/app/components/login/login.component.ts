import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  username : string = ""
  password : string = ""

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
  }

  submitForm() {
    var user = {
      username : this.username,
      password : this.password
    }
    this.authService.login(user).subscribe((token: any) => {
      this.router.navigate(['']);
      console.log(token.token);
      localStorage.setItem("token", token.token);
   })
  }
}
