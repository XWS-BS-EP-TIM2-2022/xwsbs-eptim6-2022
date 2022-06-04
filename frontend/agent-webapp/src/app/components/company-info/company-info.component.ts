import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-info',
  templateUrl: './company-info.component.html',
  styleUrls: ['./company-info.component.css']
})
export class CompanyInfoComponent implements OnInit {
  companyId! : number
  company! : any
  constructor(private route: ActivatedRoute, public service : CompanyService) { 
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
  }

  ngOnInit(): void {
    this.service.getById(this.companyId).subscribe(res => this.company = res)
  }

}
