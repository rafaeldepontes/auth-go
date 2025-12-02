package api

import (
	"github.com/rafaeldepontes/auth-go/internal/auth"
	"github.com/rafaeldepontes/auth-go/internal/middleware"
	"github.com/rafaeldepontes/auth-go/internal/user"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	UserController *user.Controller
	AuthController *auth.Controller
	Logger         *log.Logger
	Middleware     *middleware.Middleware
}
