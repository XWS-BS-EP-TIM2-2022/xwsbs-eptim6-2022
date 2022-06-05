import { Component, OnInit } from '@angular/core';
import { JobOfferService } from 'src/app/services/job-offer.service';

@Component({
  selector: 'app-job-offers',
  templateUrl: './job-offers.component.html',
  styleUrls: ['./job-offers.component.css']
})
export class JobOffersComponent implements OnInit {
  jobOffersAll : any;
  constructor(public readonly jobOfferService: JobOfferService) { }

  ngOnInit(): void {
    this.getJobOffers();
  }

  getJobOffers(){
    this.jobOffersAll = [];
    this.jobOfferService.getJobOffers().subscribe((res: any) => {
      for (let job of res){
        if ( this.convertDate(job.validTo) > new Date(new Date().toDateString()) ){
          this.jobOffersAll.push(job);
          console.log("1" + this.convertDate(job.validTo));
          console.log("2" + new Date(new Date().toDateString()))
        }
        console.log(job);
      }      
    });
  }

  convertDate(date : any) : Date{
    const date1 = new Date(date);
    return date1;
  }

}
