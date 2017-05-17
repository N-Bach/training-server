package main

import (
	"fmt"
	"log"
	"controller"
	"repo"
	"net/http"
	"encoding/json"
	
	rdb "github.com/GoRethink/gorethink"
)

func initConnection() (*rdb.Session, error) {

	session, err := rdb.Connect(rdb.ConnectOpts{
		Address: "localhost:28015",
		Database: "koyomin",
	})
	if err != nil {
		panic(err)
		return nil, err
	} 
	return session, nil
}

func fetchUser(s *rdb.Session) {
	cursor, err := rdb.DB("koyomin").Table("user").Count().Run(s)
	if err != nil {
		panic(err)
		return
	}
	var cnt int
	printStr("*** Cursor: ***")
	printObj(cursor)
	cursor.One(&cnt)
	cursor.Close()

	printStr("*** Count: ***")
	printObj(cnt)
	printStr("\n")

}

func fetchOneRecord(s *rdb.Session) {
	cursor, err := rdb.DB("koyomin").Table("user").Run(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	var person interface{}
	cursor.One(&person)
	cursor.Close()

	printStr("*** Fetch one record: ***")
	printObj(person)
	printStr("\n")
}


func main() {

	session, err := initConnection()
	if err != nil {
		panic(err)
		return
	}

	userRepo := repo.NewUserRepoRethink(session)

	// fmt.Println("connected with RethinkDb on port 28015")
	//fetchUser(session)
	//fetchOneRecord(session)

	controller := controller.NewController(userRepo)

	router := InitRoute(controller)

	fmt.Println("Server is onpen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}