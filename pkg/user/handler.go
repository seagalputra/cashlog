package user

import "github.com/gofiber/fiber/v2"

// Handler type definition for handling request/response user domain
type Handler struct {
	UserService Service
}

// Register registers a new user data
// @Summary Register a new user
// @Description Register user
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "Register user"
// @Success 201 {object} RegisterResponse{}
// @Router /v1/users/register [post]
func (h *Handler) Register(ctx *fiber.Ctx) error {
	req := new(RegisterRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.UserService.RegisterAccount(*req)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusCreated).JSON(res)

	return nil
}

// Authenticate function for user login
func (h *Handler) Authenticate(ctx *fiber.Ctx) error {
	req := new(AuthenticateRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.UserService.Authenticate(*req)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(res)

	return nil
}
