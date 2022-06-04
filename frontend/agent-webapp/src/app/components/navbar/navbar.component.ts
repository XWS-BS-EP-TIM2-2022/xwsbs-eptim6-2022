import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';
import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  company! : any
  logged: boolean = false
  role: string = ''

  constructor(public service : UsersService, public router: Router, public companyService : CompanyService) { }

  ngOnInit(): void {
    if (localStorage.getItem('jwt') != null)
      this.logged = true
    else if (localStorage.getItem('jwt') == null)
      this.logged = false

    if (localStorage.getItem('role') === 'ADMIN_ROLE')
      this.role = 'admin'
    else if (localStorage.getItem('role') === 'USER_ROLE')
      this.role = 'user'
    else if (localStorage.getItem('role') === 'COMPANY_OWNER_ROLE')
      this.role = 'owner'
    else
      this.role = ''
  }

  logout(){
    this.service.logOut()
    this.logged = false
    this.router.navigate(['/login'])
    this.role = ''
  }

  goToMyCompany(){
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
    )
  }
}
