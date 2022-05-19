import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Post } from '../model/post';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private http: HttpClient) { }

  public createNewPost(post : FormData){
    return this.http
      .post(
        '/posts/new-post',
        post, { observe: 'response', responseType: 'text' });
  }

  public commentPost(comment : Comment, id : string){
    return this.http
      .post(
        '/posts/new-comment/{id}',
        comment, { observe: 'response', responseType: 'text' });
  }

  public getAllPosts() : Observable<Post[]> {
   return this.http.get<Post[]>('/posts');
  }

  public likePost(id : string) : Observable<Post> {
    return this.http.get<Post>('/posts/like/' + encodeURIComponent(id));
  }

  public dislikePost(id : string) : Observable<Post> {
    return this.http.get<Post>('/posts/dislike/' + encodeURIComponent(id));
  }

}
