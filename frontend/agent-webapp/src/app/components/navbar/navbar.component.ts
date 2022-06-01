import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  logged: boolean = false
  role: string = ''

  constructor(public service : UsersService, public router: Router) { }

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
}
