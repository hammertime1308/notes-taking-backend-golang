package models

type User struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	SessionId string `json:"sid,omitempty"`
}
