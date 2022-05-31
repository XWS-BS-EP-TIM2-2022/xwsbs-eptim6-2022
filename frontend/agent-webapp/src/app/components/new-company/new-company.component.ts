import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-new-company',
  templateUrl: './new-company.component.html',
  styleUrls: ['./new-company.component.css']
})
export class NewCompanyComponent implements OnInit {
  errorMessage  : string = ''
  name : string = ''
  address  : string = ''
  email  : string = ''
  description  : string = ''
  
  constructor() { }

  ngOnInit(): void {
  }

}
