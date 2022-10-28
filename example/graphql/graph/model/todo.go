package model

type Tod struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"userId"`
	User *User  `json:"user"`
}
