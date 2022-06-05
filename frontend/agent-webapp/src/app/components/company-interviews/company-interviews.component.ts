import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-interviews',
  templateUrl: './company-interviews.component.html',
  styleUrls: ['./company-interviews.component.css']
})
export class CompanyInterviewsComponent implements OnInit {
  companyId!: number
  interviews!: any[]

  constructor(private route: ActivatedRoute, public service : CompanyService) { 
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
  }

  ngOnInit(): void {
    this.service.getInterviews(this.companyId).subscribe(res => this.interviews = res);
  }

}
