import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {
  rate : number = 5
  rates : number[] = [ 1, 2, 3, 4, 5]
  comment : string = ''
  salary! : number 
  difficultyLevels : string[] = ["EASY", "MEDIUM", "HARD"]
  level : string = "EASY"
  duration! : number 
  state : string = 'info'

  constructor() { }

  ngOnInit(): void {
  }

}
