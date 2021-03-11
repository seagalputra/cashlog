package user

import "github.com/gofiber/fiber/v2"

// Handler type definition for handling request/response user domain
type Handler struct {
	UserService Service
}

// Register function for registering new account
func (h *Handler) Register(ctx *fiber.Ctx) error {
	panic("Implement me")
}

// Authenticate function for user login
func (h *Handler) Authenticate(ctx *fiber.Ctx) error {
	panic("Implement me")
}
