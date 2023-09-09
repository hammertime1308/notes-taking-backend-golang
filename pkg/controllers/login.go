package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"notes-taking-backend-golang/models"
	"notes-taking-backend-golang/pkg/util"

	"github.com/sirupsen/logrus"
)

func (c *controller) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error in reading body for request. error = %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logrus.Errorf("error unmarshalling body for request. error = %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" || !util.ValidEmail(user.Email) {
		logrus.Info("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "missing body values"})
		return
	}

	user.Password = util.GetMD5Hash(user.Password)
	user.SessionId = util.GenerateSessionID()

	sessionId, err := c.repository.Login(r.Context(), user)
	if err != nil {
		logrus.Errorf("error loging user %v. error = %v", user.Email, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"sid": sessionId})
}
