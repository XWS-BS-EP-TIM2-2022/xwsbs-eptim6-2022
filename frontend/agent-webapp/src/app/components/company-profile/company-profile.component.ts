import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';

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
  exp: string[] = ["ENTRY_LEVEL", "INTERMEDIATE", "MID_LEVEL", "SENIOR"]
  experienceLevel : string = 'ENTRY_LEVEL'

  constructor(private route: ActivatedRoute, public service : CompanyService,  public matSnackBar: MatSnackBar) {
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
   }

  ngOnInit(): void {
    if (localStorage.getItem('role') === 'COMPANY_OWNER_ROLE')
      this.isOwner = false
  }

  addComment()
  {
    var comment = {
      experienceLevel : this.experienceLevel,
      text : this.comment,
      rate : this.rate,
    }

    this.service.addComment(comment, this.companyId).subscribe((data) => {
      this.matSnackBar.open("Comment successfully posted!", 'Dismiss', {
        duration: 3000
      })

      setTimeout(() => {
        window.location.reload();
      }, 1000)
    })
  }

  addSalary()
  {
    var salary = {
      experienceLevel : this.experienceLevel,
      salary : this.salary,
    }

    this.service.addSalary(salary, this.companyId).subscribe((data) => {
      this.matSnackBar.open("Salary comment successfully posted!", 'Dismiss', {
        duration: 3000
      })

      setTimeout(() => {
        window.location.reload();
      }, 1000)
    })
  }

  addInterview()
  {
    var interview = {
      experienceLevel : this.experienceLevel,
      difficultyLevel : this.level,
      selectionProcessDuration : this.duration
    }

    this.service.addInterview(interview, this.companyId).subscribe((data) => {
      this.matSnackBar.open("Interview comment successfully posted!", 'Dismiss', {
        duration: 3000
      })

      setTimeout(() => {
        window.location.reload();
      }, 1000)
    })
  }

}
