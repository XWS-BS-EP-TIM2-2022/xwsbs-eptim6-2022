import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Post, Comment } from 'src/app/model/post';
import { PostsService } from 'src/app/services/posts.service';

@Component({
  selector: 'app-posts-view',
  templateUrl: './posts-view.component.html',
  styleUrls: ['./posts-view.component.css']
})
export class PostsViewComponent implements OnInit {
  posts! : Post[];
  liked! : string[];
  disliked! : string[];
  commentText! : string;
  commentOnPost! : string;
  comments! : Comment[]

  constructor(public service : PostsService, private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.service.getAllPosts().subscribe( res => this.posts = res);
  }

  likePost(id : string) {
    this.service.likePost(id).subscribe( res => {
      this.service.getAllPosts().subscribe( res => this.posts = res);
    });
  }

  dislikePost(id : string) {
    this.service.dislikePost(id).subscribe( res => {
      this.service.getAllPosts().subscribe( res => this.posts = res);
    });
  }

  commentPost() {
    let newComment = new Comment();
   newComment = {
    Username : "",
    Text : this.commentText
   }
   this.service.commentPost(newComment, this.commentOnPost).subscribe(
    (data) => {
      this._snackBar.open('Your comment has been submited.', 'Dissmiss', {
        duration: 3000
      });

      setTimeout(() => {
      }, 1000);
    },
    (error) => {
      this._snackBar.open('Commenting failed', 'Dissmiss', {
        duration: 3000
      });
    });;;
}
}
