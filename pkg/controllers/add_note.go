package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"notes-taking-backend-golang/models"

	"github.com/sirupsen/logrus"
)

func (c *controller) addNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error in reading body for request. error = %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}

	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		logrus.Errorf("error unmarshalling body for request. error = %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}

	note := models.Note{}

	note.SessionId = data["sid"]
	note.Note = data["note"]

	if note.SessionId == "" || note.Note == "" {
		logrus.Info("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "missing body values"})
		return
	}

	latestNote, err := c.repository.AddNote(r.Context(), note)
	if err != nil {
		logrus.Errorf("error adding note for sid = %v. error = %v", note.SessionId, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"id": latestNote.Id})
}
