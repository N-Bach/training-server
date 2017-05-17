package controller

import (
	"net/http"
	"fmt"
	"html"
	"encoding/json"
	"util"
	"log"

	"entity"
	"golang.org/x/oauth2"
)

type Controller struct {
	UserRepo IUserRepo
}

func NewController(userRepo IUserRepo) *Controller{
	return &Controller{
		UserRepo: userRepo,
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (c *Controller) AuthCode(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.GoogleCode{}
	if err := decoder.Decode(&option); err != nil {
		panic(err)
	}

	fmt.Println("Code: ", option.Code)
	result := SendAuthCode(option.Code)

	json.NewEncoder(w).Encode(result.AccessToken)
}

func SendAuthCode(code string) (* oauth2.Token) {
	tok, err := util.AppConfig.Exchange(util.AppContext, code)
	if err != nil {
		log.Fatal(err)
	}
	
	result, _ := json.Marshal(tok)
 	fmt.Println(string(result))

	return tok
}
