package domain

import "time"

type User struct {
	Id			string		`json:"id"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type UserCredentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
