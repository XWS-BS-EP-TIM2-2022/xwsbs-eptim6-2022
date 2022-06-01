package mappers

import (
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"xwsbs-eptim6-2022/posts-service/store"
)

func mapToPbPost(post *store.Post) *postsServicePb.Post {
	return &postsServicePb.Post{Id: post.ID.Hex(), Username: post.Username, CreatedOn: post.CreatedOn, ImageUrl: post.ImageUrl, Text: post.Text,
		Liked: post.Liked, Disliked: post.Disliked, Comments: mapComments(post.Comments)}
}

func mapComments(comments []store.Comment) []*postsServicePb.Comment {
	var commentsResp []*postsServicePb.Comment
	for _, comment := range comments {
		commentsResp = append(commentsResp, &postsServicePb.Comment{Username: comment.Username, Text: comment.Text})
	}
	return commentsResp
}
func MapToPostResponse(post store.Post) *postsServicePb.PostResponse {
	return &postsServicePb.PostResponse{Post: mapToPbPost(&post)}
}
func MapToPostsResponse(postsDb store.Posts) *postsServicePb.PostsResponse {
	var posts []*postsServicePb.Post
	for _, post := range postsDb {
		posts = append(posts, mapToPbPost(post))
	}
	return &postsServicePb.PostsResponse{Posts: posts}
}
