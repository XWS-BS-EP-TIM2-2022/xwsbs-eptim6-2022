import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { Education, Experience, Interest, Skill, UpdateUser } from 'src/app/model/profile';
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
  qrImage: any

  constructor(public readonly profileService: ProfileService,
    private readonly formBuilder: FormBuilder, private sanitizer: DomSanitizer) {
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

  ngAfterViewInit(): void {
    this.updateUpdateForm();
  }

  updateUpdateForm() {
    this.updateFormGroup.get('name')?.setValue(this.user.name);
    this.updateFormGroup.get('surname')?.setValue(this.user.surname);
    this.updateFormGroup.get('email')?.setValue(this.user.email);
    this.updateFormGroup.get('telephone')?.setValue(this.user.telephone);
    this.updateFormGroup.get('gender')?.setValue(this.user.gender);
    this.updateFormGroup.get('birthdate')?.setValue(this.user.birthdate);
    this.updateFormGroup.get('biography')?.setValue(this.user.biography);
  }

  getUser() {
    this.profileService.getUser().subscribe((res: any) => {
      this.user = res.user
    });
    this.profileService.getUserApiKey().subscribe((res: any) => {
      this.user.apiKey = res.token
      console.log(this.user)
    });
  }
  generateApiToken() {
    this.profileService.generateApiToken().subscribe({
      next: (data) => {
        alert(data)
      },
      error: (err) => { alert(err) }
    })
    console.log("Generate api token")
  }
  update() {
    if (this.updateFormGroup.invalid) {
      alert('Invalid input');
      return;
    }

    let updateUser = new UpdateUser();
    updateUser = {
      name: this.updateFormGroup.get('name')?.value,
      surname: this.updateFormGroup.get('surname')?.value,
      email: this.updateFormGroup.get('email')?.value,
      telephone: this.updateFormGroup.get('telephone')?.value,
      gender: this.updateFormGroup.get('gender')?.value,
      birthDate: this.updateFormGroup.get('birthdate')?.value,
      biography: this.updateFormGroup.get('biography')?.value
    }

    this.profileService.updateUser(updateUser).subscribe({
      next: (data) => {
        alert("Succesfully updated!")
        this.getUser();
      },
      error: (err) => { alert("Error has occured, user not updated!") }
    });
  }

  addNewExperience() {
    if (this.experienceFormGroup.invalid) {
      alert('Invalid input');
      return;
    }

    if (this.experienceFormGroup.get('experience')?.value == "") {
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
      error: (err) => { alert("Error has occured, new experience was not added!") }
    });
  }

  addNewEducation() {
    if (this.educationFormGroup.invalid) {
      alert('Invalid input');
      return;
    }
    if (this.educationFormGroup.get('education')?.value == "") {
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
      error: (err) => { alert("Error has occured, new education was not added!") }
    });
  }

  addNewSkill() {
    if (this.skillFormGroup.invalid) {
      alert('Invalid input');
      return;
    }
    if (this.skillFormGroup.get('skill')?.value == "") {
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
      error: (err) => { alert("Error has occured, new skill was not added!") }
    });
  }

  addNewInterest() {
    if (this.interestFormGroup.invalid) {
      alert('Invalid input');
      return;
    }
    if (this.interestFormGroup.get('interest')?.value == "") {
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
      error: (err) => { alert("Error has occured, new interest was not added!") }
    });
  }

  enable2FA() {
    this.profileService.enable2FA().subscribe(data => {
      this.qrImage = this.sanitizer.bypassSecurityTrustUrl('data:image/jpeg;base64,' + data.photo)
    })
  }
}
