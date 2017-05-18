package util

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var AppContext = context.TODO()
var AppConfig = &oauth2.Config{
	ClientID:     "687936443970-lftvnktn3vcd47qp7s9jfcc3rjmjfebu.apps.googleusercontent.com",
	ClientSecret: "2YoeVzWstJk4JDCt6gWHWm4u",
	Scopes:       []string{
		"https://www.googleapis.com/auth/calendar",
	},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
	RedirectURL: "http://localhost:3000/callback",
}

var Secret = []byte("secretcode")

