package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/BurntSushi/toml"
	"github.com/eniehack/voicelog-backend/pkg/firebaseauth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	ulid "github.com/oklog/ulid/v2"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseCredential string   `toml:"firebase"`
	FrontendURL        []string `toml:"frontend"`
}

type Handler struct {
	Auth *auth.Client
}

func main() {
	configFilePath := flag.String("config", "./config.toml", "confile file's path")
	flag.Parse()

	config := new(Config)
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

	h := new(Handler)
	h.Auth = auth

	app := fiber.New()
	app.Use(cors.New(cors.Config{
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
	app.Use(firebaseauth.AuthorizationMiddleware(
		auth,
		func() bool {
			return false
		},
	))

	app.Post("/api/v1/shout/new", func(c *fiber.Ctx) error {
		uid, ok := c.Locals("uid").(string)
		if !ok {
			return c.SendStatus(http.StatusInternalServerError)
		}
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println("form")
			return c.SendStatus(http.StatusBadRequest)
		}
		file := form.File["shout"][0]
		userrec, err := auth.GetUser(c.UserContext(), uid)
		fmt.Println(userrec, file.Filename, file.Header["Content-Type"])

		/*
			switch file.Header["Content-Type"][0] {
			case "audio/ogg":
			case "audio/webm":
			case "audio/aac":
			}

			c.Response().Header.Add("Location", "")
		*/
		c.Response().Header.Add(
			"Location",
			fmt.Sprintf("https://monoxa.eniehack.net/api/v1/shout/%s", ulid.Make().String()),
		)
		return c.SendStatus(http.StatusCreated)
	})
	log.Fatal(app.Listen(":3001"))
}
