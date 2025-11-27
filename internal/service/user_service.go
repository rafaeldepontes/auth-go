package service

import (
	"net/http"

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
	// TODO: WIP
}
