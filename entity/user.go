package entity

import "time"

type User struct {
	ID        	int64     	`json:"id"`
	Email    	string    	`json:"email"`
	Name   		string     	`json:"name"`
	Password  	string    	`json:"password"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

func NewUser(email, name, password string) (*User, error){
	u := &User{
		Email: email,
		Password: password,
		Name: name,
	}
	// TODO encoder password
	return u, nil
}
