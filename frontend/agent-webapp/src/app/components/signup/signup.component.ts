import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  username: string = ""
  name: string = ""
  surname: string = ""
  email: string = ""
  phone: string = ""
  password: string = ""
  errorMessage: string = ""
  role: string = ""
  roles: string[] = ["ADMIN_ROLE", "USER_ROLE", "COMPANY_OWNER_ROLE"]

  constructor(public service: UsersService, public router : Router) { }

  ngOnInit(): void {
  }

  signup(){
    var user = {
      username : this.username,
      name : this.name,
      surname : this.surname,
      email : this.email,
      phone : this.phone,
      password : this.password,
      role : this.role
    }
    this.service.signup(user).subscribe((data) => {
      this.router.navigate(['/login']);
    },
    (error) => {
      this.errorMessage = error.error;
    })
  }

}
