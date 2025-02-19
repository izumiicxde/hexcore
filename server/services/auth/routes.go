package auth

import (
	"fmt"
	"hexcore/config"
	"hexcore/mail"
	"hexcore/types"
	"hexcore/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(s types.UserStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	router := r.Group("/auth")
	router.Post("/signup", h.Signup)                       // User Signup
	router.Post("/login", h.Login)                         // User Login
	router.Get("/verify", h.Verify)                        // Email Verification
	router.Get("/verificationCode", h.GetVerificationCode) // Request New Verification Code
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}
	if err := config.Validator.Struct(user); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}
	// Hash password before storing
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "failed to process the password"})
	}

	user.Password = hash
	user.Role = "student"

	// Email Verification Setup
	user.IsVerified = false
	user.VerificationToken = utils.GenerateVerificationCode()
	user.TokenExpiry = time.Now().Add(time.Minute * 5)

	// Store user
	if err := h.store.CreateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error creating user %v", err))
	}

	// Send verification email
	if err = mail.SendMail(user.Email, user.Register, user.VerificationToken); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error sending verification email"))
	}

	// Generate JWT token for user
	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   c.Protocol() == "https",
		SameSite: "strict",
		HTTPOnly: true,
	})
	user.Password = ""
	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "verification email sent successfully", "user": user})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := new(types.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	if err := config.Validator.Struct(req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	user, err := h.store.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	// validate the password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid credentials"))
	}
	// Generate JWT token for user
	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   c.Protocol() == "https",
		SameSite: "strict",
		HTTPOnly: true,
	})
	user.Password = ""
	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "login successfull", "user": user})
}

func (h *Handler) Verify(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request"))
	}

	// Get user from token
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
	// Get user from token
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
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("please wait before requesting another code"))
	}

	// Generate new verification code
	code := utils.GenerateVerificationCode()
	user.VerificationToken = code
	user.TokenExpiry = time.Now().Add(time.Minute * 5)
	user.IsVerified = false

	// Send email
	if err := mail.SendMail(user.Email, user.Register, code); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error sending email"))
	}
	if err := h.store.UpdateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}
	return utils.WriteJSON(c, http.StatusOK, map[string]string{"message": "verification mail sent successfully"})
}
