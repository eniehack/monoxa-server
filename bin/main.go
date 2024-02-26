package main

import (
	"context"
	"flag"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/BurntSushi/toml"
	"github.com/eniehack/voicelog-backend/internal/config"
	"github.com/eniehack/voicelog-backend/internal/handler"
	"github.com/eniehack/voicelog-backend/pkg/firebaseauth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"google.golang.org/api/option"
)

func main() {
	configFilePath := flag.String("config", "./config.toml", "confile file's path")
	flag.Parse()

	config := new(config.Config)
	if _, err := toml.DecodeFile(*configFilePath, config); err != nil {
		log.Fatalf("err reading config file: %v", err)
	}
	log.Println("loaded config file")

	ctx := context.Background()
	opt := option.WithCredentialsFile(config.FirebaseCredential)
	fb, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("firebase initialized err: %v", err)
	}
	log.Println("initialized firebase instance")
	auth, err := fb.Auth(ctx)
	if err != nil {
		log.Fatalf("firebase authentication initialized err: %v", err)
	}
	log.Println("initialized firebase authentication instance")

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("sql connection err: %v", err)
	}

	h := new(handler.Handler)
	h.Auth = auth
	h.DB = db
	h.Config = config

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginsFunc: func(origin string) bool {
			if os.Getenv("MODE") == "dev" {
				return true
			}
			for _, url := range config.FrontendURL {
				if url == origin {
					return true
				}
			}
			return false
		},
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Get("/api/v1/shout/:shout_id", h.GetShout)
	app.Get("/api/v1/notebook/:notebook_id", h.GetNotebook)

	app.Use(firebaseauth.AuthorizationMiddleware(
		auth,
		func() bool {
			return false
		},
	))

	app.Post("/api/v1/user/new", h.CreateUser)
	app.Get("/api/v1/user/me", h.Me)
	app.Post("/api/v1/shout/new", h.CreateShout)
	app.Post("/api/v1/notebook/new", h.CreateNotebook)

	log.Fatal(app.Listen(":3001"))
}
