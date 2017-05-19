package main

import (
	"controller"
	mdw "middleware"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRoute(ctrl *controller.Controller) *negroni.Negroni{
	loggingMdw := mdw.NewLoggingMiddleware()
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
	
	router.Path("/").Methods("GET").HandlerFunc(ctrl.Index)
	router.Path("/api/oauth").Methods("POST").HandlerFunc(ctrl.AuthCode)
	router.Path("/register").Methods("POST").HandlerFunc(ctrl.RegisterUser)
	router.Path("/signin").Methods("POST").HandlerFunc(ctrl.UserSignin)
	router.Path("/resource").Methods("GET").HandlerFunc(ctrl.AddLesson)

	
	router.Path("/lessons").Methods("POST").HandlerFunc(ctrl.AddLesson)
	router.Path("/lessons/enroll").Methods("POST").HandlerFunc(ctrl.AddLessonEnroll)

	
	n := negroni.New(loggingMdw, corsMdw)
	n.UseHandler(router)
	return n
}
