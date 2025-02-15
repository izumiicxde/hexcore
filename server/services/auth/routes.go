package auth

import (
	"fmt"
	"hexcore/types"
	"hexcore/utils"
	"net/http"
	"time"

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
	router.Post("/login", h.Login)
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

	//verification
	user.IsVerified = false
	user.VerificationToken = utils.GenerateVerificationCode()
	user.TokenExpiry = time.Now().Add(time.Minute * 5) // 5 minutes expiration time

	// create the user
	if err := h.store.CreateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error creating user %v", err))
	}

	// create a JWT token for the user
	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   c.Protocol() == "https",
		SameSite: "strict",
		HTTPOnly: true,
	})

	user.Password = "" // just to not send the password in the response
	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "user created successfully", "user": user})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}

	user, err := h.store.GetUserByIdentifier(req.Identifier)
	if err != nil || user == nil {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}

	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}

	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   c.Protocol() == "https",
		SameSite: "Strict",
		HTTPOnly: true,
	})

	return utils.WriteJSON(c, http.StatusOK, map[string]any{
		"message": "login successful",
		"user": map[string]any{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}
