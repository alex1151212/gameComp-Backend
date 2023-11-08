package middleware

import (
	"gameComp-Backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FileSizeLimitMiddleware(maxSizeMB int64) fiber.Handler {
	maxSizeBytes := maxSizeMB * 1024 * 1024
	return func(c *fiber.Ctx) error {
		// Check if request contains file(s)
		form, err := c.MultipartForm()
		if err != nil {
			return utils.Resp(c, fiber.StatusBadRequest, -1, nil, err.Error())
		}

		// Check each file size
		for _, files := range form.File {
			for _, file := range files {
				if file.Size > maxSizeBytes {
					return utils.Resp(c, fiber.StatusRequestEntityTooLarge, -1, nil, "File size exceeds the limit of "+strconv.FormatInt(maxSizeMB, 10)+"MB")

				}
			}
		}

		// Continue with next middleware
		return c.Next()
	}
}
