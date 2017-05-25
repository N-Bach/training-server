package main

import (
	"controller"
	mdw "middleware"
	"net/http"
	"fmt"
	"util"
	
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	jwt "github.com/dgrijalva/jwt-go"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  user := r.Context().Value("user");
  claims := user.(*jwt.Token).Claims
  fmt.Fprintf(w, "This is an authenticated request\n")
  fmt.Fprintf(w, "Claim content:\n")
  fmt.Fprintf(w, "ID: %s\n", claims.(jwt.MapClaims)["id"])
  fmt.Fprintf(w, "Email: %s\n", claims.(jwt.MapClaims)["email"])

  util.PrintObj("--------------------")
  util.PrintObj(user.(*jwt.Token))
  util.PrintObj(claims.(jwt.MapClaims))
  util.PrintObj(claims.(jwt.MapClaims)["id"])
  util.PrintObj(claims.(jwt.MapClaims)["email"])

})

func InitRoute(ctrl *controller.Controller) *negroni.Negroni{
	// loggingMdw := mdw.NewLoggingMiddleware()
	demoMdw := mdw.NewDemoMiddleware()
	corsMdw := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "POST"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "Content-Length", 
			"Accept-Encoding", "X-CSRF-Token", 
			"Authorization",
		},
	})

	router := mux.NewRouter()
	authMdw := negroni.HandlerFunc(mdw.NewJWTMiddleware(util.Secret).HandlerWithNext)
	
	
	router.Path("/").Methods("GET").HandlerFunc(ctrl.Index)
	router.Path("/api/oauth").Methods("POST").HandlerFunc(ctrl.AuthCode)

	// Public route
	router.Path("/register").Methods("POST").HandlerFunc(ctrl.RegisterUser)
	router.Path("/signin").Methods("POST").HandlerFunc(ctrl.UserSignin)
	
	// Private Route: need authentication
	router.Path("/lessons").Methods("POST").Handler(AdaptHandler(ctrl.AddLesson, authMdw))
	router.Path("/lessons/enroll").Methods("POST").Handler(AdaptHandler(ctrl.AddLessonEnroll, authMdw))
	router.Path("/feedback").Methods("POST").Handler(AdaptHandler(ctrl.AddFeedback, authMdw))
	router.Path("/review").Methods("POST").Handler(AdaptHandler(ctrl.AddReview, authMdw))

	n := negroni.New(demoMdw, corsMdw)
	n.UseHandler(router)
	return n
}

func AdaptHandler(handleFunc func(w http.ResponseWriter, r *http.Request), 
			middlewares... negroni.Handler) *negroni.Negroni{
	handler := negroni.New(middlewares...)
	handler.UseHandlerFunc(handleFunc)
	return handler
}
