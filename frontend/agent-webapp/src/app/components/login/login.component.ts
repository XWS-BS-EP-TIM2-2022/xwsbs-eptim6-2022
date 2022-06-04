import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';
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
  company! : any
  user : any

  constructor(public service : UsersService, public router : Router, public companyService : CompanyService) { }

  ngOnInit(): void {
  }

  login(){
    var user = {
      username: this.username,
      password: this.password
    }

    this.service.login(user).subscribe(
      (data) => {
        let id = localStorage.getItem('id')
        if (id == null)
          id ='0'
        this.companyService.getByOwner(id).subscribe(
          res => {
            this.company = res
            if (this.company == null)
              this.router.navigate(['/register-company'])
            else
              this.router.navigate(['/company/' + this.company.id])
          }
        );
      },
      (error) => {
        this.errorMessage = 'Invalid credentials';
        console.error('error caught');
      }
    );
  }

}
