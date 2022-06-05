import { Component, OnInit } from '@angular/core';
import { ProfileService } from 'src/app/services/profile.service';
import {MatListModule} from '@angular/material/list';


@Component({
  selector: 'app-all-users',
  templateUrl: './all-users.component.html',
  styleUrls: ['./all-users.component.css']
})
export class AllUsersComponent implements OnInit {
  users : any;
  user : any;

  constructor(public readonly profileService: ProfileService) {  }

  ngOnInit(): void {
    this.getUser();
    this.getUsers();
  }

  // ngAfterContentChecked(): void {
  //   this.getUsers();
  // }

  getUsers(){
    this.getUser();
    this.profileService.getAllUsers().subscribe((res: any) => {
      this.users = [];
      for (let u of res.users){
        if (u.username != this.user.username){
          this.users.push(u);
          console.log(u.username)
        }
      }
      console.log(this.user.username)
    });
  }

  getUser() {
    this.profileService.getUser().subscribe((res: any) => {
      this.user = res.user
    });
  }

  followUser(id : any) {
    if (id != this.user.id){
      this.profileService.followUser(id).subscribe((res: any) => {
      });
      console.log(id);
      window.location.reload();
    } 
  }

  unfollowUser(id : any){
    this.profileService.unfollowUser(id).subscribe((res: any) => {
    });
    window.location.reload();
  }

  isUserFollowing(username : any){
    let isFollowing : Boolean = false;
    for (let following of this.user.followings){
      if (username == following.username){
        isFollowing = true;
      }
    }
    return isFollowing;
  }

  isFollowingMe(username : any){
    for (let follower of this.user.followers){
      if (username == follower.username){
        return true;
      }
    }
    return false;
  }

}
