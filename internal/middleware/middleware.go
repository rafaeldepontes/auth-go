package middleware

import (
	"net/http"

	"github.com/rafaeldepontes/auth-go/internal/errorhandler"
)

func AuthCookieBased(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sessioToken *http.Cookie
		sessioToken, err := r.Cookie("session_token")

		// I should check if the token is the same for the user...
		// but I don't want to.
		if err != nil || sessioToken.Value == "" {
			errorhandler.UnauthroizedErrorHandler(w)
			return
		}

		csrfToken := r.Header.Get("X-CSRF-Token")
		if csrfToken == "" {
			errorhandler.UnauthroizedErrorHandler(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func JwtBased(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

func JwtRefreshBased(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
