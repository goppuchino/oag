package main

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
