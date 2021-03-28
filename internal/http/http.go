package http

import (
	`encoding/json`
	`net/http`
	
	`github.com/sirupsen/logrus`
)

type Service struct {
}

func New() (*Service, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return &Service{}, nil
}

func (s *Service) log(v interface{}) {
	logrus.Error(v)
}

func message(status int32, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := message(0, "Server Online")
	respond(w, response)
}