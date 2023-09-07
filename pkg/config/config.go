package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"notes-taking-backend-golang/models"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var config *models.Config

func Init(pathName string) {
	once.Do(func() {
		config = &models.Config{}
		data, err := ioutil.ReadFile(pathName)
		if err != nil {
			logrus.Fatalf("error reading json from path %v", pathName)
		}

		err = json.Unmarshal(data, config)
		if err != nil {
			logrus.Fatalf("error unmarshalling json error = %v", err.Error())
		}
	})
}

func Get() *models.Config {
	return config
}
