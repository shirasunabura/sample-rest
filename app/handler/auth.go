// Package handler ...
package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"app/usecase"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthHandler ...
type AuthHandler interface {
	Login(http.ResponseWriter, *http.Request)
}

type authHandler struct {
	authUseCase usecase.AuthUseCase
}

// NewAuthHandler ...
func NewAuthHandler(uc usecase.AuthUseCase) AuthHandler {
	return &authHandler{
		authUseCase: uc,
	}
}

// Login ...
func (h authHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)

	type request struct {
		Mail     string `json:"email"`
		Password string `json:"password"`
	}
	var t request
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if "" == t.Mail || "" == t.Password {
		log.Println("required false")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.Password = CreateHash(t.Password)
	ctx := r.Context()
	u, err := h.authUseCase.Login(ctx, t.Mail, t.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "user not found", 400)
		return
	}
	if 0 == u.ID {
		log.Println("login fail")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	type response struct {
		AccessToken string `json:"access_token"`
	}
	res := new(response)
	res.AccessToken = IssueToUserToken(u.ID)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// IssueToUserToken ...
func IssueToUserToken(uid int64) (tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "sample.com",
		"nbf": time.Now(),
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 240).Unix(),
		"uid": uid,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

// AuthorizeUserToken ...
func AuthorizeUserToken(t string) (uid int64) {
	tokenstring := strings.Split(t, " ")
	if len(tokenstring) != 2 {
		log.Println("Header Authorization Invalid Format")
		return
	}

	token, err := jwt.Parse(tokenstring[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if err != nil {
		log.Printf("jwt.Parse: %v", err)
		return
	}
	if !token.Valid {
		log.Println("token is invalid")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("token change to claims invalid")
		return
	}
	if u, ok := claims["uid"].(float64); ok {
		uid = int64(u)
	}
	return uid
}

//CreateHash Base64 パスワードのハッシュ化
func CreateHash(password string) string {

	salt := "sample"
	converted := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(converted[:])
}
