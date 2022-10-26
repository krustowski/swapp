package main

type Users struct {
	Users map[string]User `json:"users"`
}

type User struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}
