import { Component, OnInit } from '@angular/core';
import { map } from 'rxjs/operators';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-admin-requests',
  templateUrl: './admin-requests.component.html',
  styleUrls: ['./admin-requests.component.css']
})
export class AdminRequestsComponent implements OnInit {
  requests: any = []
  selectedRequest!: any

  constructor(public service : CompanyService) { }

  ngOnInit(): void {
    this.service.getAll().subscribe(res => this.requests = res);
  }

  selectRequest(request: any){
    this.selectedRequest = request;
    
  }

  approveRequest(){
    this.service.approveRequest(this.selectedRequest).subscribe();
  }
}
