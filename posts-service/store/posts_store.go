package store

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
)

type Post struct {
	ID        primitive.ObjectID `json:"id"`
	Username  string             `json:"username"`
	Text      string             `json:"text"` // slike i linkovi?
	Likes     int                `json:"likes"`
	Dislikes  int                `json:"dislikes"`
	CreatedOn string             `json:"-"`
	Comments  []Comment          `json:"comments"`
}

type Comment struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type Posts []*Post

func (p *Posts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Post) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Post) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

type PostsStore struct {
	PostsCollection *mongo.Collection
}

func (ps *PostsStore) GetAll() Posts {
	cur, err := ps.PostsCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	var posts Posts
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &elem)
	}
	cur.Close(context.TODO())
	return posts
}

func (ps *PostsStore) GetByUser(user string) Posts {
	filter := bson.D{{"username", user}}
	cur, err := ps.PostsCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	var posts Posts
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &elem)
	}
	cur.Close(context.TODO())
	return posts
}

func (ps *PostsStore) GetById(id int) Post {
	filter := bson.D{{"id", id}}
	var post Post
	err := ps.PostsCollection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}
	return post
}

func (ps *PostsStore) CreatePost(newPost Post) error {
	post, err := ps.PostsCollection.InsertOne(context.TODO(), newPost)

	if err != nil {
		return err
	}
	newPost.ID = post.InsertedID.(primitive.ObjectID)
	return nil
}

func (ps *PostsStore) LikePost(id int) Post {
	filter := bson.D{{"id", id}}

	update := bson.D{
		{"$inc", bson.D{
			{"likes", 1},
		}},
	}
	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return ps.GetById(id)
}

func (ps *PostsStore) DislikePost(id int) Post {
	filter := bson.D{{"id", id}}

	update := bson.D{
		{"$inc", bson.D{
			{"dislikes", 1},
		}},
	}
	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return ps.GetById(id)
}

func InitPostsStore() *PostsStore {
	mongoUri := "localhost:27017" //os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("posts_database").Collection("posts")
	fmt.Println(collection.Name())
	return &PostsStore{PostsCollection: collection}
}
