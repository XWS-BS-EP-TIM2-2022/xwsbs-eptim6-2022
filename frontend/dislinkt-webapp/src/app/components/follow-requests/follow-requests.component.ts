import { Component, OnInit } from '@angular/core';
import { ProfileService } from 'src/app/services/profile.service';

@Component({
  selector: 'app-follow-requests',
  templateUrl: './follow-requests.component.html',
  styleUrls: ['./follow-requests.component.css']
})
export class FollowRequestsComponent implements OnInit {
  requests : any;
  user : any;

  constructor(public readonly profileService: ProfileService) { }

  ngOnInit(): void {
    this.getUser();
  }

  getUser() {
    this.profileService.getUser().subscribe((res: any) => {
      console.log(res.user)
      this.user = res.user
    });
  }

  ConfirmRequest(username : any){
    console.log("Friend Request Confirmed!")
  }

  DeleteRequest(username : any){
    console.log("Friend Request Deleted!")
  }
}
