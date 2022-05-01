package store

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"` // slike i linkovi?
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedOn string    `json:"-"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	//ID int `json:"id"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (p *Post) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Posts []*Post

func (p *Posts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Post) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetPosts() Posts {
	return postsList
}

func CreatePost(p *Post) {
	p.ID = 3
	postsList = append(postsList, p)
}

var ErrProductNotFound = fmt.Errorf("Post not found")

func FindPost(id int) (*Post, int, error) {
	for i, p := range postsList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var postsList = []*Post{
	&Post{
		ID:        1,
		Username:  "Petra",
		Text:      "I need vacation",
		Likes:     2,
		Dislikes:  3,
		CreatedOn: time.Now().UTC().String(),
	},
	&Post{
		ID:        2,
		Username:  "Milos",
		Text:      "Good morning.",
		Likes:     3,
		Dislikes:  0,
		CreatedOn: time.Now().UTC().String(),
	},
}
