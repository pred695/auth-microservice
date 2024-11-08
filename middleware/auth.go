package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/pred695/auth-microservice/utils"
)

func VerifyUser(ctx fiber.Ctx) error {

	contextMap := fiber.Map{
		"message":    "User Verified",
		"statusText": "Token verified successfully",
	}

	token := ctx.Cookies("token")

	if token == "" {
		contextMap["statusText"] = "Bad Request"
		contextMap["message"] = "Token not found"
		return ctx.Status(fiber.StatusBadRequest).JSON(contextMap)
	}

	claims, err := utils.ValidateToken(token)
	fmt.Println(claims)
	if err != nil {
		contextMap["statusText"] = "Unauthorized"
		contextMap["message"] = "Invalid Token"
		return ctx.Status(fiber.StatusUnauthorized).JSON(contextMap)
	}

	ctx.Next()
	return nil
}
