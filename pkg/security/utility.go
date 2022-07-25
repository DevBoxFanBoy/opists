package security

import "net/http"

func GetUserNameAndPassword(r *http.Request) (string, string) {
	username, password, _ := r.BasicAuth()
	return username, password
}
