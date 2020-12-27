package libJenkins

import (
	"sort"
	"strconv"
)

func sortYearList(yearList []Job) {
	sort.Slice(yearList, func(i, j int) bool {
		year1, _ := strconv.Atoi((yearList)[i].Name)
		year2, _ := strconv.Atoi((yearList)[j].Name)
		return year1 > year2
	})
}
