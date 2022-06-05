import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-dislinkt',
  templateUrl: './company-dislinkt.component.html',
  styleUrls: ['./company-dislinkt.component.css']
})
export class CompanyDislinktComponent implements OnInit {
  companyId! : number
  company! : any
  apiConnection:any
  constructor(private route: ActivatedRoute, public service : CompanyService) { 
    this.route.params.subscribe((params) => {
      this.companyId = +params['id'];
    });
    this.apiConnection={
      apiKey:'',
      api:''
    }
  }
  ngOnInit(): void {
    this.service.getById(this.companyId).subscribe(res=>{
      if(res.apiConnection) this.apiConnection=res.apiConnection
    })
  }
  addApiConnection():void{
    console.log(this.apiConnection)
    this.service.addApiConnection(this.apiConnection,this.companyId).subscribe(res=>console.log(res))
  }

}
