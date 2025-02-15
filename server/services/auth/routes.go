package auth

import (
	"fmt"
	"hexcore/types"
	"hexcore/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(s types.UserStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/signup", h.Signup)
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	// parse the request body
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	user.Role = "student"

	// create the user
	if err := h.store.CreateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error creating user %v", err))
	}

	// create a JWT token for the user
	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		SameSite: "strict",
		HTTPOnly: true,
	})

	user.Password = "" // just to not send the password in the response
	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "user created successfully", "user": user})
}
