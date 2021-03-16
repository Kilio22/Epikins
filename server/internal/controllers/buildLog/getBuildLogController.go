package buildLog

import (
	"strconv"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/buildLog/getBuildLogService"
	"github.com/gofiber/fiber/v2"
)

func getPage(c *fiber.Ctx) (int64, internal.MyError) {
	pageString := c.Query("page")

	if pageString == "" {
		return 1, internal.MyError{}
	}
	if pageValue, err := strconv.ParseInt(pageString, 10, 64); err == nil {
		return pageValue, internal.MyError{}
	}
	return 0, internal.MyError{
		Message: "bad page query param",
		Status:  400,
	}
}

func GetBuildLogController(appData *internal.AppData, c *fiber.Ctx) error {
	page, myError := getPage(c)
	project := c.Query("project")
	starter := c.Query("starter")
	city := c.Query("city")
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	buildLogInfo, myError := getBuildLogService.GetBuildLogService(page, project, starter, city, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(buildLogInfo)
}
