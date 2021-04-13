// Package handler ...
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"app/domain/model"
	"app/usecase"
)

// UserHandler ...
type UserHandler interface {
	Index(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// NewUserHandler ...
func NewUserHandler(uc usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uc,
	}
}

// Index ...
func (h userHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Header Content-Type Not Set")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		return
	}
	if r.Header.Get("Authorization") == "" {
		log.Println("Header Authorization Not Set")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		return
	}
	uid := AuthorizeUserToken(r.Header.Get("Authorization"))
	if 0 == uid {
		log.Printf("uid can not read from jwt: %v", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
	}

	//ユーザー取得
	ctx := r.Context()
	user, _, err := h.userUseCase.FindUser(ctx, uid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if 0 == user.ID {
		log.Printf("Uesr Not Found ad : %v", r.Header.Get("Authorization"))
		http.Error(w, "Uesr Not Found", 400)
		return
	}
	type response struct {
		Name        string `json:"name"`
		NameKana    string `json:"name_kana"`
	}
	res := new(response)
	res.Name = user.FamilyName + user.FirstName
	res.NameKana = user.FamilyNameKana + user.FirstNameKana
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// Create ...
func (h userHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var t model.UserCreater
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "parameter is invalid", 400)
		log.Println(err)
		return
	}
	defer r.Body.Close()

	if "" == t.Mail || "" == t.Password {
		log.Println("required false")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "mail or pass is required", 400)
		return
	}
	t.Password = CreateHash(t.Password)
	ctx := r.Context()
	u, err := h.userUseCase.CreateUser(ctx, &t)
	if err != nil {
		log.Println(err)
		http.Error(w, "already exist", 400)
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

// Delete ...
func (h userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Header Content-Type Not Set")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		return
	}
	if r.Header.Get("Authorization") == "" {
		log.Println("Header Authorization Not Set")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		return
	}
	uid := AuthorizeUserToken(r.Header.Get("Authorization"))
	if 0 == uid {
		log.Printf("uid can not read from jwt: %v", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
	}

	ctx := r.Context()
	_, _, err := h.userUseCase.FindUser(ctx, uid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//ユーザー無効化
	err = h.userUseCase.DeleteUser(ctx, uid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.WriteHeader(204)
}
