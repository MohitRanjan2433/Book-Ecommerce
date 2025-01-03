package routes

import (
	"bookecom/controllers/otp"

	"github.com/gofiber/fiber/v2"
)

func OTPRoutes(group fiber.Router){

	group.Post("/verify-otp/:userID", otp.VerifyOTPController)
	group.Get("/resend", otp.ResendOTP)
}