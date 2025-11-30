package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rafaeldepontes/auth-go/internal/database/repository"
	"github.com/rafaeldepontes/auth-go/internal/domain"
	"github.com/rafaeldepontes/auth-go/internal/errorhandler"
	"github.com/rafaeldepontes/auth-go/internal/pagination"
	"github.com/rafaeldepontes/auth-go/internal/storage"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	userRepository *repository.UserRepository
	Logger         *log.Logger
	Cache          *storage.Caches
}

// NewUserService initialize a new UserService containing a UserRepository.
func NewUserService(userRepo *repository.UserRepository, logg *log.Logger, cache *storage.Caches) *UserService {
	return &UserService{
		userRepository: userRepo,
		Logger:         logg,
		Cache:          cache,
	}
}

// FindAllUsers list all the users without a filter and returns each
// one with pagination and a few datas missing for LGPD.
func (s *UserService) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infoln("Listing all the users in the database...")

	sizeStr := r.URL.Query().Get("size")
	if sizeStr == "" {
		sizeStr = "25"
	}
	size, _ := strconv.Atoi(sizeStr)

	currentPageStr := r.URL.Query().Get("page")
	if currentPageStr == "" {
		currentPageStr = "1"
	}
	currentPage, _ := strconv.Atoi(currentPageStr)

	var users []repository.User
	users, totalRecords, err := s.userRepository.FindAllUsers(size, currentPage)
	if err != nil {
		errorhandler.BadRequestErrorHandler(w, err, r.URL.Path)
		s.Logger.Errorf("An error occurred: %v", err)
		return
	}

	pageModel := pagination.NewPagination(users, uint(currentPage), uint(totalRecords), uint(size))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	s.Logger.Infof("Found %v users from a total of %v", len(users), totalRecords)

	json.NewEncoder(w).Encode(pageModel)
}

// FindUserById list an user by his id and returns a none
// pagination result and a few datas missing for LGPD.
func (s *UserService) FindUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	s.Logger.Infof("Listing user by id - %v", idStr)

	if idStr == "" {
		errorhandler.BadRequestErrorHandler(w, errorhandler.ErrIdIsRequired, r.URL.Path)
		s.Logger.Errorf("An error occurred: %v", errorhandler.ErrIdIsRequired)
		return
	}

	pathId, _ := strconv.Atoi(idStr)
	id := uint(pathId)

	var user *repository.User
	user, err := s.userRepository.FindUserById(id)
	if err != nil {
		errorhandler.BadRequestErrorHandler(w, errorhandler.ErrUserNotFound, r.URL.Path)
		s.Logger.Errorf("An error occurred: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (s *UserService) FindUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	s.Logger.Infof("Listing user by username: %v", username)

	if username == "" {
		errorhandler.BadRequestErrorHandler(w, errorhandler.ErrUsernameIsRequired, r.URL.Path)
		s.Logger.Errorf("An error occurred: %v", errorhandler.ErrUsernameIsRequired)
		return
	}

	var user *repository.User
	user, err := s.userRepository.FindUserByUsername(username)
	if err != nil {
		errorhandler.BadRequestErrorHandler(w, errorhandler.ErrUserNotFound, r.URL.Path)
		s.Logger.Errorf("An error occurred: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

// UpdateUserDetails changes the user age and/or name if it's the account owner.
func (s *UserService) UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infoln("Updating an user")

	var user *repository.User
	username := r.PathValue("username")
	user, err := s.userRepository.FindUserByUsername(username)
	if err != nil {
		s.Logger.Errorf("An error occurred: %v", err)
		errorhandler.BadRequestErrorHandler(w, errorhandler.ErrUserNotFound, r.URL.Path)
		return
	}

	var newUserDetails domain.UserDetails
	if err := json.NewDecoder(r.Body).Decode(&newUserDetails); err != nil {
		s.Logger.Errorf("An error occurred: %v", err)
		errorhandler.InternalErrorHandler(w)
		return
	}

	if err := isValidUserDetails(user, &newUserDetails); err != nil {
		s.Logger.Errorf("An error occurred: %v", err)
		errorhandler.BadRequestErrorHandler(w, err, r.URL.Path)
		return
	}

	user.Age = &newUserDetails.Age

	err = s.userRepository.UpdateUserDetails(user)
	if err != nil {
		s.Logger.Errorf("An error occurred: %v", err)
		errorhandler.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteAccount deletes the user from the database by his username
// if it's the account owner.
func (s *UserService) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	s.Logger.Infof("Deleting an account by his username: %v\n", username)

	err := s.userRepository.DeleteAccount(username)
	if err != nil {
		s.Logger.Errorf("An error occurred: %v", err)
		errorhandler.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Account deleted successfully")
}

func isValidUserDetails(user *repository.User, userRequest *domain.UserDetails) error {
	if userRequest.Age <= 0 {
		return errorhandler.ErrAgeIsRequired
	}

	if userRequest.Age == *user.Age {
		return errorhandler.ErrEqualAge
	}

	return nil
}
