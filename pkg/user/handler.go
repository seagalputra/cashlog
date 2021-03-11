package user

import "github.com/gofiber/fiber/v2"

// Handler type definition for handling request/response user domain
type Handler struct {
	UserService Service
}

// Register function for registering new account
func (h *Handler) Register(ctx *fiber.Ctx) error {
	request := new(RegisterRequest)

	errResponse := struct {
		IsError bool   `json:"is_error"`
		Message string `json:"message"`
		Data    interface{}
	}{
		IsError: true,
		Message: "Failed to register user",
		Data:    nil,
	}

	err := ctx.BodyParser(&request)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	userID, err := h.UserService.RegisterAccount(request)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	successResponse := struct {
		IsError bool   `json:"is_error"`
		Message string `json:"message"`
		Data    interface{}
	}{
		IsError: false,
		Message: "Success",
		Data:    userID,
	}
	ctx.Status(fiber.StatusOK).JSON(successResponse)

	return nil
}

// Authenticate function for user login
func (h *Handler) Authenticate(ctx *fiber.Ctx) error {
	panic("Implement me")
}
