import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { of } from 'rxjs';
import { CompanyService } from 'src/app/services/company.service';
import { JobOfferService } from 'src/app/services/job-offer.service';

@Component({
  selector: 'app-new-job-offer',
  templateUrl: './new-job-offer.component.html',
  styleUrls: ['./new-job-offer.component.css']
})
export class NewJobOfferComponent implements OnInit {
  company! : any
  errorMessage : string = ''
  position : string = ''
  exp: string[] = ["ENTRY_LEVEL", "INTERMEDIATE", "MID_LEVEL", "SENIOR"]
  experience: string = ''
  shareOnDislinkt: boolean = false
  description: string = ''
  workSchedule : string = ''
  hoursPerWeek : number = 0
  validTo = new FormControl(new Date())

  constructor(public service : JobOfferService, public matSnackBar: MatSnackBar, public companyService : CompanyService) { }

  ngOnInit(): void {
  }

  createJobOffer() {
    var offer = {
      position : this.position,
      shareOnDislinkt : this.shareOnDislinkt,
      validTo : this.validTo.value,
      description : this.description,
      experience : this.experience,
      workSchedule : this.workSchedule,
      hoursPerWeek : this.hoursPerWeek
    }

    this.getCompanyId()
    if (this.company.status === 'PENDING') {
      this.matSnackBar.open("Your company still needs to be verified!", 'Dismiss', {
        duration: 2000
      })

      setTimeout(() => {
        window.location.reload();
      }, 1000)
    }else{
      this.service.createJobOffer(offer, this.company.id).subscribe((data) => {
        this.matSnackBar.open("Job offer successfully posted!", 'Dismiss', {
          duration: 2000
        })
  
        setTimeout(() => {
          window.location.reload();
        }, 1000)
      },
      (error) => {
        this.errorMessage = error.error;
      })
    }
  }

  getCompanyId() {
    let id = localStorage.getItem('id')
    if (id == null)
      id ='0'
    this.companyService.getByOwner(id).subscribe(
      res => {
        this.company = res
      }
    );
  }

}
