package firebaseauth

import (
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware(auth *auth.Client, isAuthorizeNotNeeded func() bool) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if isAuthorizeNotNeeded() {
			return c.Next()
		}
		authHeader := string(c.Request().Header.Peek("Authorization")[:])
		if len(authHeader) < 1 || strings.HasPrefix(authHeader, "Bearer") == false {
			return c.SendStatus(http.StatusBadRequest)
		}

		idToken := strings.Split(authHeader, " ")[1]
		token, err := auth.VerifyIDToken(c.UserContext(), idToken)
		if err != nil {
			log.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		c.Locals("uid", token.UID)
		return c.Next()
	}
}
