package domain

import "time"

type Session struct {
	Id           string
	Username     string
	IsRevoked    bool
	RefreshToken string
	CreatedAt    time.Time
	ExpiresAt    time.Time
}
