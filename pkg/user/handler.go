package user

import "github.com/gofiber/fiber/v2"

// Handler type definition for handling request/response user domain
type Handler struct {
	UserService Service
}

// Register function for registering new account
func (h *Handler) Register(ctx *fiber.Ctx) error {
	request := new(RegisterRequest)

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	response, err := h.UserService.RegisterAccount(request)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(response)

	return nil
}

// Authenticate function for user login
func (h *Handler) Authenticate(ctx *fiber.Ctx) error {
	panic("Implement me")
}
