package buildLog

import (
	"net/http"
	"strconv"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/buildLog/exportBuildLogService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func getTimestampFromQuery(stringValue string) (int64, internal.MyError) {
	finalValue, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return 0, util.GetMyError("invalid query param", err, http.StatusBadRequest)
	}
	return finalValue, internal.MyError{}
}

func ExportBuildLogController(appData *internal.AppData, c *fiber.Ctx) error {
	start, myError := getTimestampFromQuery(c.Query("start"))
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	end, myError := getTimestampFromQuery(c.Query("end"))
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	project := c.Query("project")
	city := c.Query("city")
	bytes, myError := exportBuildLogService.ExportBuildLogService(start, end, project, city, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	c.Set(fiber.HeaderContentType, "text/csv; charset=utf-8")
	c.Status(http.StatusOK)
	return c.Send(bytes)
}
