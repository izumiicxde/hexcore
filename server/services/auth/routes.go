package auth

import (
	"fmt"
	"hexcore/mail"
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

	router.Get("/verify", h.Verify)
	router.Get("/verificationCode", h.GetVerificationCode)
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

	if err = mail.SendMail(user.Email, user.Username, user.VerificationToken); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error sending verification email, please try again later"))
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
	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "verification email sent successfully", "user": user})
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

func (h *Handler) Verify(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request"))
	}

	// get the jwt token from cookies
	_, claims, err := utils.ParseJWT(c.Cookies("token"))
	if err != nil {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("invalid token"))
	}

	userId := claims["userId"].(float64)
	user, err := h.store.GetUserById(uint(userId))
	if err != nil {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("no user found"))
	}
	if user.IsVerified {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("email already verified"))
	}

	if user.VerificationToken != code || user.TokenExpiry.Before(time.Now()) {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("invalid verification code"))
	}

	user.IsVerified = true
	user.TokenExpiry = time.Now()
	user.VerificationToken = "-"

	if err := h.store.UpdateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error updating user %v", err))
	}
	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "email verified successfully", "redirect": true})
}

func (h *Handler) GetVerificationCode(c *fiber.Ctx) error {
	// get the jwt token from cookies
	_, claims, err := utils.ParseJWT(c.Cookies("token"))
	if err != nil {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("invalid token"))
	}

	userId := claims["userId"].(float64)
	user, err := h.store.GetUserById(uint(userId))
	if err != nil {
		return utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("no user found"))
	}
	if user.IsVerified || time.Now().Before(user.TokenExpiry) {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("please wait sometime, before requesting the code"))
	}

	code := utils.GenerateVerificationCode()
	user.VerificationToken = code
	user.TokenExpiry = time.Now().Add(time.Minute * 5)
	user.IsVerified = false

	if err := mail.SendMail(user.Email, user.Username, code); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("there was an error sending the email %v", err))
	}
	if err := h.store.UpdateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}
	return utils.WriteJSON(c, http.StatusOK, map[string]string{"message": "verification mail send successfully"})
}
