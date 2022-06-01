import { Component, OnInit } from '@angular/core';
import { CompanyService } from 'src/app/services/company.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-new-company',
  templateUrl: './new-company.component.html',
  styleUrls: ['./new-company.component.css']
})
export class NewCompanyComponent implements OnInit {
  errorMessage  : string = ''
  name : string = ''
  address  : string = ''
  email  : string = ''
  phone  : string = ''
  description  : string = ''
  culture  : string = ''
  yearOfEstablishment : number = 0
  
  constructor(public service : CompanyService,  public matSnackBar: MatSnackBar) { }

  ngOnInit(): void {
  }

  registerCompany()
  {
    var company = {
      name : this.name,
      address : this.address,
      email : this.email,
      phone : this.phone,
      description : this.description,
      culture : this.culture,
      yearOfEstablishment : this.yearOfEstablishment
    }
    this.service.registerNewCompany(company).subscribe(
      (data) => {
        this.matSnackBar.open("Request successfully sent!", 'Dismiss', {
          duration: 2000
        })

        setTimeout(() => {
          window.location.reload();
        }, 1000)
      },
      (error) => {
        this.errorMessage = error.error;
      }
    )

  }

}
