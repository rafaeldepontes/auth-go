package user

import "net/http"

type Service interface {
	FindAllUsersCursorPagination(w http.ResponseWriter, r *http.Request)
	FindAllUsersOffSetPagination(w http.ResponseWriter, r *http.Request)
	FindUserById(w http.ResponseWriter, r *http.Request)
	FindUserByUsername(w http.ResponseWriter, r *http.Request)
	UpdateUserDetails(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}