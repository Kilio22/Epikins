package utils

import (
	"strconv"

	"github.com/gofiber/fiber"
)

func GetQueryBoolValue(queryKey string, c *fiber.Ctx) (bool, error) {
	shouldUpdateString := c.Query(queryKey)
	var shouldUpdateProjectList bool

	if shouldUpdateString != "" {
		shouldUpdateValue, err := strconv.ParseBool(shouldUpdateString)
		if err != nil {
			return false, err
		} else {
			shouldUpdateProjectList = shouldUpdateValue
		}
	}
	return shouldUpdateProjectList, nil
}
