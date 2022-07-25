package security

import (
	"github.com/casbin/casbin/v2"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func checkPermission(e *casbin.Enforcer, r *http.Request) (bool, error) {
	user, _ := GetUserNameAndPassword(r)
	method := r.Method
	path := r.URL.Path
	result, err := e.Enforce(user, path, method)
	log.Println("Check Permission for", user, method, path, result, err)
	return result, err
}

func Authorizer(e *casbin.Enforcer) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if ok, _ := checkPermission(e, r); ok {
			next(w, r)
		} else {
			http.Error(w, http.StatusText(403), 403)
		}
	}
}
