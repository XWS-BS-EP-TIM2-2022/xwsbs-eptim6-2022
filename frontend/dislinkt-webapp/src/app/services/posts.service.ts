import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Comment, Post } from '../model/post';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private http: HttpClient) { }

  public createNewPost(post : FormData){
    return this.http
      .post(
        '/api/posts',
        post, { observe: 'response', responseType: 'text' });
  }

  public commentPost(comment : Comment, id : string){
    return this.http
      .post(
        '/api/posts/'+ encodeURIComponent(id) + '/comments',
        comment, { observe: 'response', responseType: 'text' });
  }

  public getAllPosts() : Observable<any> {
   return this.http.get<any>('/api/posts');
  }

  public likePost(id : string) {
    return this.http.put('/api/posts/'  + encodeURIComponent(id)+ '/likes', id);
  }

  public dislikePost(id : string){
    return this.http.put('/api/posts/'  + encodeURIComponent(id)+ '/dislikes', id);
  }

}
