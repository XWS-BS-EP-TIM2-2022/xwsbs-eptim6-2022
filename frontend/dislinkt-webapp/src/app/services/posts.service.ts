import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private http: HttpClient) { }

  public createNewPost(post : FormData){
    return this.http
      .post(
        'http://localhost:9090/posts/new-post',
        post, { observe: 'response', responseType: 'text' });
  }

}
