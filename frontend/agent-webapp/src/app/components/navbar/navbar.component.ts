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

  constructor(public service : UsersService, public router: Router) { }

  ngOnInit(): void {
    if (localStorage.getItem('jwt') != null)
      this.logged = true
    else
      this.logged = false
  }

  logout(){
    this.service.logOut()
    this.logged = false
    this.router.navigate(['/login'])
  }
}
