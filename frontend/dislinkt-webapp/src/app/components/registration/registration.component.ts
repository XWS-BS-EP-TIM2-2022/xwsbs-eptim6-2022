import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  username : string = "";
  email : string = "";
  password : string = "";
  repeatedPassword : string = "";
  name : string = "";
  surname : string = "";
  differentPasswords : boolean = false;


  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
  }

  submitForm() {
    if (this.password === this.repeatedPassword){
      this.differentPasswords = false;
      var user = {
        username : this.username,
        password : this.password,
        name : this.name,
        surname : this.surname,
        email : this.email
      }
      this.authService.register(user).subscribe((token: any) => {
        this.router.navigate(['']);
      })
    }
    else {
      this.differentPasswords = true;
    }
  }

}
