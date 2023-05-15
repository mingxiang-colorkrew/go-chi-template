package middleware

import (
	"measure/config"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

func NewAuthMiddleware(app *config.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			token, claims, err := jwtauth.FromContext(r.Context())

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if token == nil || jwt.Validate(token) != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// TODO: actually use user ID in claims to check user from DB / cache / other providers
			app.Logger().Sugar().Infow("Request is authorized", "claims", claims)

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
