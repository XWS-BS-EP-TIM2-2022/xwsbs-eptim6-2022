package handlers

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"strconv"
	"time"
)

type ImagesHandler struct {
	bucket *storage.BucketHandle
}

func InitImageHandler() (*ImagesHandler, error) {
	config := &firebase.Config{
		StorageBucket: "xml-proj.appspot.com",
	}
	sa := option.WithCredentialsFile("./token.json")
	app, err := firebase.NewApp(context.Background(), config, sa)
	if err != nil {
		return nil, err
	}
	client, err := app.Storage(context.Background())
	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}
	return &ImagesHandler{bucket: bucket}, nil
}

func (ih *ImagesHandler) SaveImage(image []byte) (url string, err error) {
	name := generateImageName() //TODO: Mozda dodati i username korisnika koji objavljuje radi dodatne sigurnosti
	writer := ih.bucket.Object(name).NewWriter(context.TODO())
	_, err = writer.Write(image)
	err = writer.Close()
	imageUrl := generateUrl(name)
	fmt.Println(imageUrl)
	return imageUrl, err
}

func generateUrl(name string) string {
	baseFirebaseUrl := "https://firebasestorage.googleapis.com/v0/b/xml-proj.appspot.com/o/"
	return baseFirebaseUrl + name + "?alt=media"
}

func generateImageName() string {
	name := time.Now().Unix()
	strconv.FormatInt(name, 16)
	return strconv.FormatInt(name, 16)
}
