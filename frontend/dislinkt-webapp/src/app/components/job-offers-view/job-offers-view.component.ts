import { Component, OnInit } from '@angular/core';
import { JobOffersService } from 'src/app/services/job-offers.service';

@Component({
  selector: 'app-job-offers-view',
  templateUrl: './job-offers-view.component.html',
  styleUrls: ['./job-offers-view.component.css']
})
export class JobOffersViewComponent implements OnInit {

  jobOffersAll : any;
  constructor(public readonly jobOfferService: JobOffersService) { }

  ngOnInit(): void {
    this.getJobOffers();
  }

  getJobOffers(){
    //this.jobOffersAll = [];
    this.jobOfferService.getJobOffers().subscribe((res: any) => {
      console.log(res);
      //this.jobOffersAll = res;
      for (let job of res.offers){
        if ( this.convertDate(job.validTo) > new Date(new Date().toDateString()) ){
          this.jobOffersAll.push(job);
        }
      }      
    });
  }

  convertDate(date : any) : Date{
    const date1 = new Date(date);
    return date1;
  }

}
