package api

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaeldepontes/auth-go/internal/database"
	"github.com/rafaeldepontes/auth-go/internal/repository"
	"github.com/rafaeldepontes/auth-go/internal/service"
)

// Init initialize all the resources needed for the server run properly
func Init() (*Configuration, *Application, *sql.DB, error) {
	godotenv.Load(".env", ".env.example")

	config := &Configuration{
		JwtBasedPort:        os.Getenv("JWT_PORT"),
		CookieBasedPort:     os.Getenv("COOKIE_PORT"),
		JwtRefreshBasedPort: os.Getenv("JWT_REFRESH_PORT"),
	}

	db, err := database.Open()

	var userRepository *repository.UserRepository = repository.NewUserRepository(db)
	var userService *service.UserService = service.NewUserService(userRepository)

	application := &Application{
		UserService: userService,
	}

	return config, application, db, err
}
