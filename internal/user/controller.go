package user

import "net/http"

type Controller interface {
	ListAllCursor(w http.ResponseWriter, r *http.Request)
	ListAllOffset(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	UpdateDetails(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}