import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';
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

  constructor(public service : PostsService) { }

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

}
