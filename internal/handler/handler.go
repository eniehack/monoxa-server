package handler

import (
	"database/sql"

	"firebase.google.com/go/v4/auth"
	"github.com/eniehack/voicelog-backend/internal/config"
)

type Handler struct {
	Auth   *auth.Client
	DB     *sql.DB
	Config *config.Config
}
