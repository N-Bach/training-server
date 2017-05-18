package main

import (
	"fmt"
	"log"
	"controller"
	"repo"
	"net/http"
	
	rdb "github.com/GoRethink/gorethink"
)

func main() {

	session, err := rdb.Connect(rdb.ConnectOpts{
		Address: "localhost:28015",
		Database: "hikaru",
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("connected with RethinkDb on port 28015")
	//fetchUser(session)
	//fetchOneRecord(session)

	userRepo := repo.NewUserRepoRethink(session)
	lessonRepo := repo.NewLessonRepoRethink(session)

	controller := controller.NewController(userRepo, lessonRepo)

	router := InitRoute(controller)

	fmt.Println("Server is onpen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
