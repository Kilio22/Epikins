package utils

import (
	"epikins-api/pkg/libJenkins"
	"errors"

	"github.com/gofiber/fiber"
)

func GetVisibilityQueryParam(c *fiber.Ctx) (libJenkins.Visibility, error) {
	visibilityString := c.Query("visibility")
	visibility := libJenkins.PRIVATE

	if visibilityString != "" {
		_, ok := libJenkins.VisibilityMap[visibilityString]
		if !ok {
			return visibility, errors.New("invalid visibility parameter")
		}
		visibility = libJenkins.VisibilityMap[visibilityString]
	}
	return visibility, nil
}
