package errorhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	ErrInvalidUsernameOrPassword = errors.New("Error: invalid username or password")
	ErrInvalidTokenSigningMethod = errors.New("Error: invalid token signing method")
	ErrInvalidCursorSignature    = errors.New("Error: cursor signature is invalid")
	ErrInvalidTokenSignature     = errors.New("Error: token signature is invalid")
	ErrInvalidExpiredToken       = errors.New("Error: invalid token, token already expired")
	ErrInvalidCursorLength       = errors.New("Error: invalid cursor length")
	ErrUsernameIsRequired        = errors.New("Error: username is required")
	ErrPasswordIsRequired        = errors.New("Error: password is required")
	ErrInvalidTokenClaim         = errors.New("Error: invalid token claims")
	ErrUserAlreadyExists         = errors.New("Error: user already exist")
	ErrTokenNotValidYet          = errors.New("Error: token is not valid yet")
	ErrInvalidCSRFToken          = errors.New("Error: CSRF token missing")
	ErrSessionNotFound           = errors.New("Error: session not found")
	ErrMalformedToken            = errors.New("Error: token is malformed")
	ErrAgeIsRequired             = errors.New("Error: age is required")
	ErrInvalidMethod             = errors.New("Error: invalid method")
	ErrCreatingToken             = errors.New("Error: error while creating token")
	ErrEqualUsername             = errors.New("Error: the new username should be different from the actual")
	ErrIdIsRequired              = errors.New("Error: identifier is required")
	ErrUserNotFound              = errors.New("Error: user not found")
	ErrParsingToken              = errors.New("Error: error parsing token")
	ErrInvalidToken              = errors.New("Error: token missing or invalid")
	ErrTokenRevoked              = errors.New("Error: session token revoked")
	ErrInvalidType               = errors.New("Error: unsupported type")
	ErrInvalidId                 = errors.New("Error: invalid username, needs to be your own")
	ErrEqualAge                  = errors.New("Error: the new age should be different from the actual")
)

const BrazilianDateTimeFormat = "02/01/2006 15:04:05"

type Error struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path,omitempty"`
	Status    int    `json:"status"`
}

var (
	BadRequestErrorHandler = func(w http.ResponseWriter, err error, path string) {
		writeError(w, err.Error(), http.StatusBadRequest, path)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Occurred.", http.StatusInternalServerError, "")
	}
	UnauthroizedErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusUnauthorized, "")
	}
	ForbiddenErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusForbidden, "")
	}
	RequestErrorHandler = func(w http.ResponseWriter, err error, status int, path string) {
		writeError(w, err.Error(), status, path)
	}
)

func writeError(w http.ResponseWriter, message string, status int, path string) {
	var timestamp string = time.Now().Format(BrazilianDateTimeFormat)
	resp := Error{
		Status:    status,
		Message:   message,
		Path:      path,
		Timestamp: timestamp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error(err)
	}
}
