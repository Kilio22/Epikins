package controllers

import (
	"strconv"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/getBuildLogsService"
	"github.com/gofiber/fiber/v2"
)

func getPage(c *fiber.Ctx) (int64, internal.MyError) {
	pageString := c.Query("page")

	if pageString == "" {
		return 1, internal.MyError{}
	}
	if pageValue, err := strconv.ParseInt(pageString, 10, 64); err != nil {
		return pageValue, internal.MyError{}
	}
	return 0, internal.MyError{
		Message: "bad page query param",
		Status:  400,
	}
}

func GetBuildLogsController(appData *internal.AppData, c *fiber.Ctx) error {
	page, myError := getPage(c)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	buildLogsInfo, myError := getBuildLogsService.GetBuildLogsService(page, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(buildLogsInfo)
}
