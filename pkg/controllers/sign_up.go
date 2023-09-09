package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"notes-taking-backend-golang/models"
	"notes-taking-backend-golang/pkg/util"

	"github.com/sirupsen/logrus"
)

func (c *controller) signUp(w http.ResponseWriter, r *http.Request) {
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

	if user.Name == "" || user.Email == "" || user.Password == "" || !util.ValidEmail(user.Email) {
		logrus.Info("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "missing body values"})
		return
	}

	user.SessionId = util.GenerateSessionID()
	user.Password = util.GetMD5Hash(user.Password)

	err = c.repository.AddNewUser(r.Context(), user)
	if err != nil {
		logrus.Errorf("error adding new user. error = %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "user added successfully"})
}
