import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-comments',
  templateUrl: './company-comments.component.html',
  styleUrls: ['./company-comments.component.css']
})
export class CompanyCommentsComponent implements OnInit {
  companyId!: number
  comments!: any[]

  constructor(private route: ActivatedRoute, public service : CompanyService) {
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
   }

  ngOnInit(): void {
    this.service.getComments(this.companyId).subscribe(res => this.comments = res);
  }

}