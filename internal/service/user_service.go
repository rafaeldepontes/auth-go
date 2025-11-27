package service

import (
	"encoding/json"
	"net/http"

	"github.com/rafaeldepontes/auth-go/internal/errorhandler"
	"github.com/rafaeldepontes/auth-go/internal/repository"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	Logger         *log.Logger
	userRepository *repository.UserRepository
}

// NewUserService initialize a new UserService containing a UserRepository
func NewUserService(userRepo *repository.UserRepository, logg *log.Logger) *UserService {
	return &UserService{
		Logger:         logg,
		userRepository: userRepo,
	}
}

// FindAllUsers list all the users without a filter and returns each
// one with pagination and a few datas missing for LGPD
func (us *UserService) FindAllUsers(w http.ResponseWriter, r *http.Request) {

	us.Logger.Infoln("[LOG] Listing all the users in the database...")
	var users []repository.User
	users, err := us.userRepository.FindAllUsers()
	if err != nil {
		errorhandler.BadRequestErrorHandler(w, err, r.URL.Path)
		log.Errorf("[ERROR] An error occurred: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	us.Logger.Infof("[LOG] Found %v users", len(users))

	//TODO: IMPLEMENT PAGINATION FOR THIS ENDPOINT
	json.NewEncoder(w).Encode(users)
}
