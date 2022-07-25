package security

import (
	"context"
	"github.com/DevBoxFanBoy/opists/pkg/config"
	"net/http"
)

type AuthenticationMiddleWare struct {
	config config.Config
}

func NewAuthenticationMiddleWare(config config.Config) *AuthenticationMiddleWare {
	return &AuthenticationMiddleWare{config: config}
}

func (a *AuthenticationMiddleWare) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth, password := GetUserNameAndPassword(r)
	rw.Header().Set("WWW-Authenticate", `Basic realm="Restricted OPISTS"`)
	if (auth == "" && password == "") || a.config.Opists.Security.AdminUsername == "" ||
		a.config.Opists.Security.AdminPassword == "" {
		http.Error(rw, http.StatusText(401), 401)
		return
	}
	if auth == a.config.Opists.Security.AdminUsername && password == a.config.Opists.Security.AdminPassword {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		next(rw, r.WithContext(context.WithValue(ctx, CurrentUser{}, auth)))
	} else {
		users := a.config.Opists.Security.Users
		for _, user := range users {
			if auth == user.Username && password == user.Password {
				ctx := r.Context()
				if ctx == nil {
					ctx = context.Background()
				}
				next(rw, r.WithContext(context.WithValue(ctx, CurrentUser{}, auth)))
				return
			}
		}
		http.Error(rw, http.StatusText(401), 401)
		return
	}
}
