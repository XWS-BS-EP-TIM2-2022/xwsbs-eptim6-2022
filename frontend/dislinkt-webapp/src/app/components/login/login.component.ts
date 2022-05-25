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
  errorMessage : string = ""
  logged : boolean = false

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
    if (localStorage.getItem('jwt') != null)
      this.logged = true
    else
      this.logged = false
  }

  submitForm() {
    var user = {
      username : this.username,
      password : this.password
    }
    this.authService.login(user).subscribe(
      (data) => {
        this.router.navigate(['/homepage'])
      },
      (error) => {
        this.errorMessage = 'Invalid credentials';
        console.error('error caught');
      }
    );
  }
}
