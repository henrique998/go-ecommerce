package createaccountcontroller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

func (cc *createAccountController) Handle(c fiber.Ctx) error {
	req := requests.CreateAccountRequest{
		Name:  "jhon doe",
		Email: "jhondoe@gmail.com",
		Pass:  "123456",
	}

	cc.us.Execute(req)

	return nil
}