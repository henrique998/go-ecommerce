package loginservice

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

func (ctrl *loginController) Handle(c fiber.Ctx) error {
	body := c.Body()

	var req requests.LoginRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	accessToken, refreshToken, err := ctrl.service.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	accessTokenCookie := fiber.Cookie{
		Name:     "goauth:access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Path:     "/",
	}

	refreshTokenCookie := fiber.Cookie{
		Name:     "goauth:refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		Path:     "/",
	}

	c.Cookie(&accessTokenCookie)
	c.Cookie(&refreshTokenCookie)

	return nil
}
