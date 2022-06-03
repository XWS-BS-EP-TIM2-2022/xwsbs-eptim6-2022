import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Education, Experience, Interest, Skill } from 'src/app/model/profile';
import { ProfileService } from 'src/app/services/profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  user: any;
  id: number = 0;
  followBtn: any;
  followersNum: number = 0;
  followingNum: number = 0;
  public readonly updateFormGroup: FormGroup;
  public readonly experienceFormGroup: FormGroup;
  public readonly educationFormGroup: FormGroup;
  public readonly skillFormGroup: FormGroup;
  public readonly interestFormGroup: FormGroup;

  constructor(public readonly profileService: ProfileService,
              private readonly formBuilder: FormBuilder) { 
                this.updateFormGroup = this.formBuilder.group({
                  id: [],
                  email: ['', Validators.compose([Validators.required, Validators.email])],
                  name: [],
                  surname: [],
                  telephone: [],
                  gender: [],
                  birthdate: [],
                  biography: []
                });
                this.experienceFormGroup = this.formBuilder.group({
                  experience: ''
                });
                this.educationFormGroup = this.formBuilder.group({
                  education: ''
                });
                this.skillFormGroup = this.formBuilder.group({
                  skill: ''
                });
                this.interestFormGroup = this.formBuilder.group({
                  interest: ''
                });
              }

  ngOnInit(): void {
    this.getUser();
  }

  getUser() {
    this.profileService.getUser().subscribe((res: any) => {
      console.log(res.user)
      this.user = res.user
    });
  }

  update() {
    if (this.updateFormGroup.invalid) {
      alert('Invalid input');
      return;
    }

    this.profileService.updateUser(this.updateFormGroup.getRawValue()).subscribe({
      next: (data) => {
      alert("Succesfully updated!")
      this.getUser();
    },
      error: (err) => {alert("Error has occured, user not updated!")}
    });
  }

  addNewExperience() {
    if (this.experienceFormGroup.invalid) {
      alert('Invalid input');
      return;
    }

    if (this.experienceFormGroup.get('experience')?.value == ""){
      alert('invalid input');
      return;
    }

    let newExperience = new Experience;
    newExperience = {
      text: this.experienceFormGroup.get('experience')?.value
    }

    this.profileService.addNewExperience(newExperience).subscribe({
      next: (data) => {
      alert("Succesfully added!")
      this.getUser();
    },
      error: (err) => {alert("Error has occured, new experience was not added!")}
    });
  }

  addNewEducation() {
    if (this.educationFormGroup.invalid) {
      alert('Invalid input');
      return;
    }
    if (this.educationFormGroup.get('education')?.value == ""){
      alert('invalid input');
      return;
    }

    let newEducation = new Education();
    newEducation = {
      text: this.educationFormGroup.get('education')?.value
    }

    this.profileService.addNewEducation(newEducation).subscribe({
      next: (data) => {
      alert("Succesfully added!")
      this.getUser();
    },
      error: (err) => {alert("Error has occured, new education was not added!")}
    });
  }

  addNewSkill() {
    if (this.skillFormGroup.invalid) {
      alert('Invalid input');
      return;
    }
    if (this.skillFormGroup.get('skill')?.value == ""){
      alert('invalid input');
      return;
    }

    let newSkill = new Skill();
    newSkill = {
      text: this.skillFormGroup.get('skill')?.value
    }

    this.profileService.addNewSkill(newSkill).subscribe({
      next: (data) => {
      alert("Succesfully added!")
      this.getUser();
    },
      error: (err) => {alert("Error has occured, new skill was not added!")}
    });
  }

  addNewInterest() {
    if (this.interestFormGroup.invalid) {
      alert('Invalid input');
      return;
    }
    if (this.interestFormGroup.get('interest')?.value == ""){
      alert('invalid input');
      return;
    }

    let newInterest = new Interest();
    newInterest = {
      text: this.interestFormGroup.get('interest')?.value
    }

    this.profileService.addNewInterest(newInterest).subscribe({
      next: (data) => {
      alert("Succesfully added!")
      this.getUser();
    },
      error: (err) => {alert("Error has occured, new interest was not added!")}
    });
  }

}
