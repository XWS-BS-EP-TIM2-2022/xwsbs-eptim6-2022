import { Component, OnInit } from '@angular/core';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-all-companies',
  templateUrl: './all-companies.component.html',
  styleUrls: ['./all-companies.component.css']
})
export class AllCompaniesComponent implements OnInit {
  companies! : any[]

  constructor(public service : CompanyService) { }

  ngOnInit(): void {
    this.service.getAll().subscribe(res => this.companies = res);
  }

}
