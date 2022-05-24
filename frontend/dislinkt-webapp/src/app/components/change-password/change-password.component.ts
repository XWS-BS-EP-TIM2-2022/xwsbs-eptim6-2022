import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {
  oldPassword : string = ''
  newPassword : string = ''
  repeatedPassword : string = ''
  errorMessage : string = ''

  constructor() { }

  ngOnInit(): void {
  }

  submitForm(){
    if (this.isFormDataValid()){
      this.errorMessage = ""
      this.oldPassword = ""
      this.newPassword = ""
      this.repeatedPassword = ""
    }
     
  }

  isFormDataValid() : boolean {
    var passwordPattern = new RegExp(/(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*\W)/);
    if (this.newPassword !== this.repeatedPassword)
    {
      this.errorMessage = "Passwords do not match";
      return false;
    }
    if (this.newPassword.length < 8)
    {
      this.errorMessage = "Password must be minimum 8 characters long";
      return false;
    }
    if (!passwordPattern.test(this.newPassword)){
      this.errorMessage = "Incorrect password format.";
      return false;
    }
    return true;
  }
}
