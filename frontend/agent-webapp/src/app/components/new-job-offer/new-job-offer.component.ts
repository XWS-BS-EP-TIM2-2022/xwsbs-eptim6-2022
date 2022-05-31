import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-new-job-offer',
  templateUrl: './new-job-offer.component.html',
  styleUrls: ['./new-job-offer.component.css']
})
export class NewJobOfferComponent implements OnInit {
  errorMessage : string = ''
  position : string = ''
  constructor() { }

  ngOnInit(): void {
  }

}
