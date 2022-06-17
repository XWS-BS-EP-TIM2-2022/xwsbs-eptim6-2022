import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
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
  logged: boolean = false
  TwoFA!: boolean;
  code: string = ""

  constructor(private authService: AuthService, private router: Router, public matSnackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.passwordless = false;
    if (localStorage.getItem('jwt') != null)
      this.logged = true
    else
      this.logged = false

    this.TwoFA = false;
  }

  login() {
    this.authService.findUserByUsername(this.username).subscribe(
      data => {
        if (data.user.MFAStatus)
          this.TwoFA = true;
        else
          this.submitForm();
      },
      error => {
        this.errorMessage = "Invalid credentials"
      });
  }

  submitForm() {
    var user = {
      username: this.username,
      password: this.password
    }
    this.authService.login(user).subscribe(
      (data) => {
        window.location.href = "/home"
      },
      (error) => {
        this.errorMessage = error.error.message;
        console.error('error caught');
      }
    );
  }

  passwordlessLogin() {
    this.authService.passwordless(this.email).subscribe(
      (data) => {
        this.passwordless = false;
        this.matSnackBar.open("Email successfully sent!", 'Dismiss', {
          duration: 2000
        })

        setTimeout(() => {
          window.location.reload();
        }, 1000)
      },
      (error) => {
        this.errorMessage = error.error;
      }
    )
  }

  submit2FA() {
    this.authService.submit2FA(this.code).subscribe(data => {
      this.submitForm();
    },
      error => {
        this.errorMessage = error.error
      }
    )
  }
}
