package controllers

import (
	"encoding/json"
	"net/http"
	"notes-taking-backend-golang/models"

	"github.com/sirupsen/logrus"
)

func (c *controller) deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logrus.Errorf("error unmarshalling body for request. error = %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}

	if body["sid"] == "" || body["id"] == "" {
		logrus.Info("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "missing body values"})
		return
	}

	note := models.Note{
		Id:        body["id"],
		SessionId: body["sid"],
	}
	err = c.repository.DeleteNote(r.Context(), note)
	if err != nil {
		logrus.Errorf("error deleting note for session id = %v, note id = %v. error = %v", note.SessionId, note.Id, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "note deleted"})
}
