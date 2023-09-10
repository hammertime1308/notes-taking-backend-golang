package controllers

import (
	"encoding/json"
	"net/http"
	"notes-taking-backend-golang/models"

	"github.com/sirupsen/logrus"
)

func (c *controller) getAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	user.SessionId = r.URL.Query().Get("sid")

	if user.SessionId == "" {
		logrus.Info("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "missing session id"})
		return
	}

	notes, err := c.repository.GetAllNotes(r.Context(), user)
	if err != nil {
		logrus.Errorf("error fetching all notes for session = %v. error = %v", user.SessionId, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"notes": notes})
}
