package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"xwsbs-eptim6-2022/posts-service/handlers"
)

func main() {
	l := log.New(os.Stdout, "posts-service ", log.LstdFlags)
	ph := handlers.NewPostsHandler(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/posts", ph.GetAll)
	//getRouter.HandleFunc("/posts/homepage", ph.GetForUser)
	//getRouter.HandleFunc("/posts/user-profile", ph.GetByUser)
	getRouter.HandleFunc("/posts/{id:[0-9]+}", ph.GetOne)
	//getRouter.HandleFunc("/posts/like/{id:[0-9]+}", ph.LikePost)
	//getRouter.HandleFunc("/posts/dislike/{id:[0-9]+}", ph.DislikePost)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/posts/create", ph.CreatePost)
	//postRouter.HandleFunc("/posts/comment/{id:[0-9]+}", ph.CommentPost)

	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
