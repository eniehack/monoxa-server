package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/eniehack/voicelog-backend/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

func (h *Handler) GetNotebook(c *fiber.Ctx) error {
	q := database.New(h.DB)
	shout, err := q.GetNotebook(c.UserContext(), c.Params("notebook_id"))
	if errors.Is(err, sql.ErrNoRows) {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"meta": map[string]string{
			"name":     shout.Name,
			"shout_id": shout.AliasID,
		},
		"shouts": []string{
			"a",
		},
	})
}

type CreateNotebookRequestParams struct {
	Name       string `json:"name"`
	NotebookID string `json:"notebook_id"`
}

func (h *Handler) CreateNotebook(c *fiber.Ctx) error {
	uid, ok := c.Locals("uid").(string)
	if !ok {
		return c.SendStatus(http.StatusInternalServerError)
	}
	param := new(CreateNotebookRequestParams)
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("form")
		return c.SendStatus(http.StatusBadRequest)
	}
	file := form.File["shout"][0]
	userrec, err := h.Auth.GetUser(c.UserContext(), uid)
	fmt.Println(userrec, file.Filename, file.Header["Content-Type"])

	q := database.New(h.DB)
	user, err := q.GetUserID(c.UserContext(), uid)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	tx, err := h.DB.Begin(c.UserContext())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	q = database.New(h.DB)
	qtx := q.WithTx(tx)
	tenantUlid := ulid.Make().String()
	createTenantParam := database.CreateNotebookParams{
		Ulid:    tenantUlid,
		AliasID: param.NotebookID,
		Name:    param.Name,
	}
	if err := qtx.CreateNotebook(c.UserContext(), createTenantParam); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	userParam := database.AddUserToNotebookParams{
		NotebookID: tenantUlid,
		UserID:     user.Ulid,
	}
	if err := qtx.AddUserToNotebook(c.UserContext(), userParam); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if err := tx.Commit(c.UserContext()); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	c.Response().Header.Add(
		"Location",
		fmt.Sprintf("%s/api/v1/notebook/%s", h.Config.FrontendURL[0], param.NotebookID),
	)
	return c.SendStatus(http.StatusCreated)
}
