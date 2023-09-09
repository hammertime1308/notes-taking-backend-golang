package models

type Note struct {
	Id        string `json:"id"`
	Note      string `json:"note"`
	SessionId string
}
