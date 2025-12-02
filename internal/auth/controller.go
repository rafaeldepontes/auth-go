package auth

import "net/http"

type Controller interface {
	RegisterEp(w http.ResponseWriter, r *http.Request)
	LoginCookieBasedEp(w http.ResponseWriter, r *http.Request)
	LoginJwtBasedEp(w http.ResponseWriter, r *http.Request)
	LoginJwtRefreshBasedEp(w http.ResponseWriter, r *http.Request)
	RenewAccessTokenEp(w http.ResponseWriter, r *http.Request)
	RevokeSessionEp(w http.ResponseWriter, r *http.Request)
	GetAuthCallbackOAuth2Ep(w http.ResponseWriter, r *http.Request)
	LogoutOAuth2Ep(w http.ResponseWriter, r *http.Request)
	GetAuthOAuth2Ep(w http.ResponseWriter, r *http.Request)
}
