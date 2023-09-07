package models

type Config struct {
	ServeOn  int      `json:"serveOn"`
	Database database `json:"database"`
	LogLevel string   `json:"logLevel"`
}

type database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}
