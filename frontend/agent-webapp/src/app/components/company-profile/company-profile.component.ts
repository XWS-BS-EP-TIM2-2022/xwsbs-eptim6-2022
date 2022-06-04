import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {
  companyId!: number
  isOwner : boolean = false

  rate : number = 5
  rates : number[] = [ 1, 2, 3, 4, 5]
  comment : string = ''
  salary! : number 
  difficultyLevels : string[] = ["EASY", "MEDIUM", "HARD"]
  level : string = "EASY"
  duration! : number 
  state : string = 'info'

  constructor(private route: ActivatedRoute) {
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
   }

  ngOnInit(): void {
    if (localStorage.getItem('role') === 'COMPANY_OWNER_ROLE')
      this.isOwner = true
  }

}
