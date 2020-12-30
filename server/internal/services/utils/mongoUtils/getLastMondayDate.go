package mongoUtils

import "time"

func GetLastMondayDate() int64 {
	var lastMonday time.Time

	currentDay := time.Now().Weekday()
	if currentDay == time.Sunday {
		lastMonday = time.Now().AddDate(0, 0, -6)
	} else {
		lastMonday = time.Now().AddDate(0, 0, -int(currentDay)+1)
	}
	year, month, day := lastMonday.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, lastMonday.Location()).Unix()
}
