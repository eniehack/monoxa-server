package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/eniehack/voicelog-backend/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

type CreateUserRequestParams struct {
	Name    string `json:"name"`
	AliasID string `json:"user_id"`
}

type Notebook struct {
	NotebookID string `json:"notebook_id"`
}

func (h *Handler) Me(c *fiber.Ctx) error {
	uid, ok := c.Locals("uid").(string)
	if !ok {
		log.Println("cannot serialize uid to string")
		return c.SendStatus(http.StatusInternalServerError)
	}
	q := database.New(h.DB)
	user, err := q.GetUserID(c.UserContext(), uid)
	if err == sql.ErrNoRows {
		return c.SendStatus(http.StatusNotFound)
	} else if err != nil {
		log.Printf("handler/Me err: %v\n", err)
		return c.SendStatus(http.StatusBadRequest)
	}
	notebooksRow, err := q.GetBelongNotebook(c.UserContext(), uid)
	if err != nil && err != sql.ErrNoRows {
		return c.SendStatus(http.StatusInternalServerError)
	}
	var notebooks []Notebook
	for _, notebook := range notebooksRow {
		notebooks = append(notebooks, Notebook{NotebookID: notebook.AliasID})
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"user_id":   user.AliasID,
		"name":      user.Name,
		"notebooks": notebooks,
	})
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	uid, ok := c.Locals("uid").(string)
	if !ok {
		return c.SendStatus(http.StatusInternalServerError)
	}
	body := new(CreateUserRequestParams)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	tx, err := h.DB.Begin(c.UserContext())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	q := database.New(h.DB)
	userUlid := ulid.Make().String()
	notebookUlid := ulid.Make().String()
	txq := q.WithTx(tx)
	userParams := database.CreateUserParams{
		Ulid:    userUlid,
		Uid:     uid,
		AliasID: body.AliasID,
		Name:    body.Name,
	}
	txq.CreateUser(c.UserContext(), userParams)
	noteParams := database.CreateNotebookParams{
		Ulid:    notebookUlid,
		AliasID: body.AliasID,
		Name:    body.Name,
	}
	txq.CreateNotebook(c.UserContext(), noteParams)
	paticipantsParams := database.AddUserToNotebookAsAdminParams{
		NotebookID: notebookUlid,
		UserID:     userUlid,
	}
	txq.AddUserToNotebookAsAdmin(c.UserContext(), paticipantsParams)
	if err := tx.Commit(c.UserContext()); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.SendStatus(http.StatusCreated)
}
