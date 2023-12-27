package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pworld/go-generator/internal/user/models/entity"
)

// FormatUserDetails formats the user data for the response
func FormatUserDetails(user entity.User) map[string]interface{} {
	return map[string]interface{}{
		// Fields to format here...
		"id":        user.ID,
		"fullname":  user.Fullname,
		"email":     user.Email,
		"phone":     user.Phone,
		"username":  user.Username,
		"password":  user.Password,
		"createdat": user.CreatedAt,
		"updatedat": user.UpdatedAt,
	}
}

// UserResponse formats a successful user-related response
func UserResponse(c *fiber.Ctx, user entity.User) error {
	formattedUser := FormatUserDetails(user)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    formattedUser,
	})
}

// UserListResponse formats a response for a list of users
func UserListResponse(c *fiber.Ctx, users []entity.User, total, totalPages int) error {
	formattedUsers := make([]map[string]interface{}, 0)
	for _, user := range users {
		formattedUsers = append(formattedUsers, FormatUserDetails(user))
	}
	return c.JSON(fiber.Map{
		"success":    true,
		"items":      formattedUsers,
		"total":      total,
		"totalPages": totalPages,
	})
}

// UserErrorResponse formats an error response specific to user operations
func UserErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   message,
		"context": "user operation",
	})
}
