package security

import (
	"log"
	"net/http"
)

type CurrentUser struct{}

func LogCurrentUser(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Current User", r.Context().Value(CurrentUser{}).(string))
	next(rw, r)
}
