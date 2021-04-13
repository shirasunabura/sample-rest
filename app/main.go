// main
package main

import (
	"app/infrastructure/dao"
	"fmt"
	"log"
	"net/http"
	"os"

	"app/handler"
	"app/infrastructure/persistence"
	"app/usecase"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func main() {
	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNINGKEY")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r := mux.NewRouter()
	s := r.PathPrefix("/v1").Subrouter()

	conn := &dao.Connection{}
	conn.Connect()
	defer conn.CloseConn()

	userPersistence := persistence.NewUserPersistence(conn)
	authPersistence := persistence.NewAuthPersistence(conn)

	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)
	s.HandleFunc("/user", userHandler.Create).Methods("POST")
	s.Handle("/user", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(userHandler.Index)))).Methods("GET")
	s.Handle("/user", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(userHandler.Delete)))).Methods("PUT")

	authUseCase := usecase.NewAuthUseCase(authPersistence)
	authHandler := handler.NewAuthHandler(authUseCase)
	s.HandleFunc("/auth", authHandler.Login).Methods("POST")
	
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
