package http

import (
	`encoding/json`
	`net/http`
	
	`travel/internal/user`
)

func(s *Service) Auth (w http.ResponseWriter, r *http.Request) {
	
	account := &user.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		respond(w, message(1, "Invalid request"))
		return
	}
	
	resp, err := account.Login()
	if err != nil {
		s.log(err)
		respond(w, message(1, "not login"))
	}
	
	respond(w, resp)
}
