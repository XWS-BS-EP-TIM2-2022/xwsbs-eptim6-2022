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

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
  }

  submitForm() {
    var user = {
      username : this.username,
      password : this.password
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
}
