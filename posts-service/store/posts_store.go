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

type User struct {
	Username string `json:"username"`
}

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Text      string             `bson:"text"`
	Likes     int                `bson:"likes"`
	Liked     []string           `bson:"liked"`
	Dislikes  int                `bson:"dislikes"`
	Disliked  []string           `bson:"disliked"`
	CreatedOn string             `bson:"createdOn"`
	ImageUrl  string             `bson:"url"`
	Comments  []Comment          `bson:"comments"`
}

type Comment struct {
	Username string `bson:"username"`
	Text     string `bson:"text"`
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

func (ps *PostsStore) GetAll() (Posts, error) {
	cur, err := ps.PostsCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	var posts Posts
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &elem)
	}
	cur.Close(context.TODO())
	return posts, nil
}

func (ps *PostsStore) GetByUser(user string) (Posts, error) {
	filter := bson.D{{"username", user}}
	cur, err := ps.PostsCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return nil, err
	}
	var posts Posts
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &elem)
	}
	cur.Close(context.TODO())
	return posts, nil
}

func (ps *PostsStore) GetById(id primitive.ObjectID) (Post, error) {
	filter := bson.D{{"_id", id}}
	var post Post
	err := ps.PostsCollection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (ps *PostsStore) CreatePost(newPost Post) (error, primitive.ObjectID) {
	post, err := ps.PostsCollection.InsertOne(context.TODO(), newPost)
	newPost.ID = post.InsertedID.(primitive.ObjectID)
	if err != nil {
		return err, newPost.ID
	}
	return nil, newPost.ID
}

func (ps *PostsStore) LikePost(id primitive.ObjectID, user string) (Post, error) {
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$inc", bson.D{
			{"likes", 1},
		}},
		{"$push", bson.D{
			{"liked", user},
		}},
	}
	var post Post
	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return post, err
	}

	return ps.GetById(id)
}

func (ps *PostsStore) UnlikePost(id primitive.ObjectID, user string) (Post, error) {
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$inc", bson.D{
			{"likes", -1},
		}},
		{"$pull", bson.D{
			{"liked", user},
		}},
	}
	var post Post
	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return post, err
	}

	return ps.GetById(id)
}

func (ps *PostsStore) DislikePost(id primitive.ObjectID, user string) (Post, error) {
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$inc", bson.D{
			{"dislikes", 1},
		}},
		{"$push", bson.D{
			{"disliked", user},
		}},
	}
	var post Post
	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return post, err
	}

	return ps.GetById(id)
}

func (ps *PostsStore) UndislikePost(id primitive.ObjectID, user string) (Post, error) {
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$inc", bson.D{
			{"dislikes", -1},
		}},
		{"$pull", bson.D{
			{"disliked", user},
		}},
	}
	var post Post
	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return post, err
	}

	return ps.GetById(id)
}

func (ps *PostsStore) InsertComment(id primitive.ObjectID, comment Comment) error {
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$push", bson.D{
			{"comments", comment},
		}},
	}

	_, err := ps.PostsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ps *PostsStore) IsAlreadyLiked(id primitive.ObjectID, username string) bool {
	post, _ := ps.GetById(id)
	for _, element := range post.Liked {
		if element == username {
			return true
		}
	}
	return false
}

func (ps *PostsStore) IsAlreadyDisliked(id primitive.ObjectID, username string) bool {
	post, _ := ps.GetById(id)
	for _, element := range post.Disliked {
		if element == username {
			return true
		}
	}
	return false
}

func (ps *PostsStore) Drop() {
	ps.Drop()
}

func InitPostsStore(mongoUri string) *PostsStore {
	//os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	//collection := client.Database("posts_database").Collection("user_posts") drop?
	collection := client.Database("post_db").Collection("user_posts")
	fmt.Println(collection.Name())
	return &PostsStore{PostsCollection: collection}
}
