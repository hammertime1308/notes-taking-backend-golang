package models

type Note struct {
	Id        string `json:"id" db:"id"`
	Note      string `json:"note" db:"note"`
	SessionId string `db:"session_id" json:"-"`
}
