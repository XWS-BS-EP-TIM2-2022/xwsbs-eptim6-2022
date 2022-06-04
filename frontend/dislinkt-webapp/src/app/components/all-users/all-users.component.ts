import { Component, OnInit } from '@angular/core';
import { ProfileService } from 'src/app/services/profile.service';

@Component({
  selector: 'app-all-users',
  templateUrl: './all-users.component.html',
  styleUrls: ['./all-users.component.css']
})
export class AllUsersComponent implements OnInit {
  users : any;
  user : any;

  constructor(public readonly profileService: ProfileService) { this.getUser(); }

  ngOnInit(): void {
    this.getUsers();
  }

  // ngAfterContentChecked(): void {
  //   this.getUsers();
  // }

  getUsers(){
    this.profileService.getAllUsers().subscribe((res: any) => {
      this.users = res.users;
      // for (let u of res.users) {
      //   if (u.username != this.user.username){
          // console.log(u.username)
          // const index = res.users.indexOf(u);
          // if (index > -1) {
          //   this.users.splice(index, 1); // 2nd parameter means remove one item only
          // }
      //     this.users.push(u);
      //   }
      // }
      console.log(this.users)
    });
  }

  getUser() {
    this.profileService.getUser().subscribe((res: any) => {
      console.log(res.user)
      this.user = res.user
    });
  }

  followUser(username : any) {
    if (username != this.user.username){
      this.profileService.followUser().subscribe((res: any) => {
      });
      console.log(username);
    } 
  }

  isUserFollowing(username : any){
    let isFollowing : Boolean = false;
    for (let following of this.user.followings){
      console.log(following)
      if (username == following.username){
        isFollowing = true;
      }
    }
    return isFollowing;
  }

}
