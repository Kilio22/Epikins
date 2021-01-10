package controllerUtil

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetQueryBoolValue(queryKey string, defaultValue bool, c *fiber.Ctx) (bool, error) {
	valueString := c.Query(queryKey)
	var value = defaultValue

	if valueString != "" {
		parsedValue, err := strconv.ParseBool(valueString)
		if err != nil {
			return false, err
		} else {
			value = parsedValue
		}
	}
	return value, nil
}
