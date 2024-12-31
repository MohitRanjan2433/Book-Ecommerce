package otp

import (
	"bookecom/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func VerifyOTPController(c *fiber.Ctx) error {
	userIDstr := c.Params("userID")
	otp := c.FormValue("otp")

	if otp == "" {
		return fiber.NewError(fiber.StatusBadRequest, "OTP is required")
	}

	userID, err := uuid.Parse(userIDstr)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid user ID format",
		})
	}

	err = service.VerifyOTP(userID, otp)
	if err != nil{
		log.Printf("Error verifying OTP: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": "Failed to verify OTP",
		})
	}

	// Return a success response if OTP is verified successfully
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "OTP verified successfully",
	})
}