package mappers

import (
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"xwsbs-eptim6-2022/posts-service/store"
)

func MapToPostResponse(post store.Post) *postsServicePb.PostResponse {
	return &postsServicePb.PostResponse{Post: &postsServicePb.Post{Username: post.Username, CreatedOn: post.CreatedOn, ImageUrl: post.ImageUrl, Text: post.Text}}
}
