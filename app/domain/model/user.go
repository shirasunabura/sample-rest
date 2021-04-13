// Package model ...
package model

import "time"

// User ...
type User struct {
	ID             int64
	Mail           string
	Password       string
	FamilyName     string
	FirstName      string
	FamilyNameKana string
	FirstNameKana  string
	LastLoginAt    time.Time
}

// UserCreater ...
type UserCreater struct {
	FamilyName     string `json:"family_name"`
	FirstName      string `json:"first_name"`
	FamilyNameKana string `json:"family_name_kana"`
	FirstNameKana  string `json:"first_name_kana"`
	Post           int32  `json:"post"`
	Prefecture     string `json:"prefecture"`
	City           string `json:"city"`
	Block          string `json:"block"`
	HouseNumber    string `json:"house_number"`
	Number         string `json:"number"`
	Building       string `json:"building"`
	Phone          string `json:"phone"`
	Mail           string `json:"email"`
	Password       string `json:"password"`
}
