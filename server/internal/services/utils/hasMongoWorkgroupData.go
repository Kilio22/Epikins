package utils

import "epikins-api/internal"

func HasMongoWorkgroupData(groupName string, workgroupsData []internal.MongoWorkgroupData) (internal.MongoWorkgroupData, bool) {
	for _, workgroupData := range workgroupsData {
		if groupName == workgroupData.Name {
			return workgroupData, true
		}
	}
	return internal.MongoWorkgroupData{}, false
}
