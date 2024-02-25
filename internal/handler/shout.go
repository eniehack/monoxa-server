package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/eniehack/voicelog-backend/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

func (h *Handler) GetShout(c *fiber.Ctx) error {
	q := database.New(h.DB)
	shout, err := q.GetShout(c.UserContext(), c.Params("shout_id"))
	if errors.Is(err, sql.ErrNoRows) {
		return c.SendStatus(http.StatusNotFound)
	}
	c.Status(http.StatusOK)
	return c.JSON(map[string]string{
		"file": shout.Ulid,
		"user": shout.UserID,
	})
}

func (h *Handler) CreateShout(c *fiber.Ctx) error {
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
	userrec, err := h.Auth.GetUser(c.UserContext(), uid)
	fmt.Println(userrec, file.Filename, file.Header["Content-Type"])
	shout_id := ulid.Make().String()

	tenantIDBuf := new(strings.Builder)
	tenantIDFile, err := form.File["tenant"][0].Open()
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	io.Copy(tenantIDBuf, tenantIDFile)
	q := database.New(h.DB)
	user, err := q.GetUserID(c.UserContext(), uid)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	param := database.CreateShoutParams{
		UserID:     user.Ulid,
		Ulid:       shout_id,
		NotebookID: tenantIDBuf.String(),
	}
	q.CreateShout(c.UserContext(), param)

	/*
		switch file.Header["Content-Type"][0] {
		case "audio/ogg":
		case "audio/webm":
		case "audio/aac":
		}

		c.Response().Header.Add("Location", "")
	*/
	basename := filepath.Base(file.Filename)

	if err := c.SaveFile(
		file,
		filepath.Join(
			h.Config.FileBucket,
			fmt.Sprintf("%s%s", shout_id, filepath.Ext(basename)),
		),
	); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	c.Response().Header.Add(
		"Location",
		fmt.Sprintf("https://monoxa.eniehack.net/api/v1/shout/%s", shout_id),
	)
	return c.SendStatus(http.StatusCreated)
}
