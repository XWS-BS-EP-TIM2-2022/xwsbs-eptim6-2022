import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  username: string = ""
  password: string = ""
  errorMessage: string = ""
  passwordless!: boolean
  email!: string
  checkMail = new FormControl('', [Validators.email]);

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
    this.passwordless = false;
  }

  submitForm() {
    var user = {
      username: this.username,
      password: this.password
    }
    this.authService.login(user).subscribe(
      (data) => {
        this.router.navigate(['/change-password'])
      },
      (error) => {
        this.errorMessage = 'Invalid credentials';
        console.error('error caught');
      }
    );
  }

  passwordlessLogin() {
    this.authService.passwordless(this.email).subscribe(
      (data) => {
        this.passwordless = false;
        window.location.reload;
      },
      (error) => {
        this.errorMessage = error.error;
      }
    )
  }
}
