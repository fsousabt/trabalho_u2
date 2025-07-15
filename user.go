package main

type User struct {
	Login    string `json:"login"`
	UserId   int    `json:"userid"`
	Region   int    `json:"region"`
	Password string `json:"password"`
}
