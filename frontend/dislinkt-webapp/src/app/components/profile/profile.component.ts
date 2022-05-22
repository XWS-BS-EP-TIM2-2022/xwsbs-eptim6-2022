import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProfileService } from 'src/app/services/profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  user: any;
  followBtn: any;
  followersNum: number = 0;
  followingNum: number = 0;
  public readonly myFormGroup: FormGroup;

  constructor(public readonly profileService: ProfileService,
              private readonly formBuilder: FormBuilder) { 
                this.myFormGroup = this.formBuilder.group({
                  id: [],
                  email: ['', Validators.compose([Validators.required, Validators.email])],
                  name: [],
                  surname: [],
                  telephone: [],
                  gender: [],
                  birthdate: [],
                  biography: []
              });
              }

  ngOnInit(): void {
  }

  update() {

  }

  addNewExperience() {

  }

  addNewEducation() {
    
  }

  addNewSkill() {
    
  }

  addNewInterest() {
    
  }

}
