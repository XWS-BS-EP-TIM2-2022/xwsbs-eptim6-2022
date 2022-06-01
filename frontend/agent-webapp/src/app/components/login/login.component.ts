import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  username: string = ""
  password: string = ""
  errorMessage: string = ""

  constructor(public service : UsersService, public router : Router) { }

  ngOnInit(): void {
  }

  login(){
    var user = {
      username: this.username,
      password: this.password
    }

    this.service.login(user).subscribe(
      (data) => {
        this.router.navigate(['/register-company'])
      },
      (error) => {
        this.errorMessage = 'Invalid credentials';
        console.error('error caught');
      }
    );
  }

}
