package auth

import "net/http"

type Service interface {
	Register(w http.ResponseWriter, r *http.Request)
	LoginCookieBased(w http.ResponseWriter, r *http.Request)
	LoginJwtBased(w http.ResponseWriter, r *http.Request)
	LoginJwtRefreshBased(w http.ResponseWriter, r *http.Request)
	RenewAccessToken(w http.ResponseWriter, r *http.Request)
	RevokeSession(w http.ResponseWriter, r *http.Request)
	GetAuthCallbackOAuth2(w http.ResponseWriter, r *http.Request)
	LogoutOAuth2(w http.ResponseWriter, r *http.Request)
	GetAuthOAuth2(w http.ResponseWriter, r *http.Request)
}
