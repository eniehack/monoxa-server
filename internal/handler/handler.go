package handler

import (
	"firebase.google.com/go/v4/auth"
	"github.com/eniehack/voicelog-backend/internal/config"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	Auth   *auth.Client
	DB     *pgx.Conn
	Config *config.Config
}
