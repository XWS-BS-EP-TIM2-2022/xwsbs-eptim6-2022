package handlers

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"image"
	"image/jpeg"
	"log"
	"net/http"
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

func (ih *ImagesHandler) SaveImageHandler(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println("ERROR")
	}
	f, _, err := r.FormFile("file")
	imageData, _, err := image.Decode(f)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, imageData, nil)
	if err != nil {
		log.Fatalln("ENCOIDNG IMAGE ERROR")
	}
	send_s3 := buf.Bytes()
	url, err := ih.SaveImage(send_s3)
	if err != nil {
		fmt.Println("Greska")
	}
	defer f.Close()
	json.NewEncoder(rw).Encode(url)
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
	//name = strings.Join(strings.Split(name, " "), "")
	//name = strings.Split(name, "+")[0]
	return strconv.FormatInt(name, 16)
}
