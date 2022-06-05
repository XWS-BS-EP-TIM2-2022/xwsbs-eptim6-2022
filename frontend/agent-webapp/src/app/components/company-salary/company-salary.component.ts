import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-salary',
  templateUrl: './company-salary.component.html',
  styleUrls: ['./company-salary.component.css']
})
export class CompanySalaryComponent implements OnInit {
  companyId!: number
  salaries!: any[]


  constructor(private route: ActivatedRoute, public service : CompanyService) {
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
   }

  ngOnInit(): void {
    this.service.getSalaries(this.companyId).subscribe(res => this.salaries = res);
  }

}
