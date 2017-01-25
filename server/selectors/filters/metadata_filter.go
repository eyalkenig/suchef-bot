package filters

import (
	"github.com/eyalkenig/suchef-bot/server/models"
)

func SelectMetadata(courses []*models.Course, metadataTypeName string, metadataTypeValue *string) []*models.Course {
	filteredCourses := []*models.Course{}
	for _, course := range courses {
		courseMetadataValues := course.Tags[metadataTypeName]
		if courseMetadataValues != nil {
			for _, metadataValue := range courseMetadataValues {
				if *metadataValue == *metadataTypeValue {
					filteredCourses = append(filteredCourses, course)
				}
			}
		}
	}

	return filteredCourses
}

func ExcludeMetadata(courses []*models.Course, metadataTypeName string, metadataTypeValue *string) []*models.Course {
	filteredCourses := []*models.Course{}
	for _, course := range courses {
		courseMetadataValues := course.Tags[metadataTypeName]
		if courseMetadataValues == nil {
			filteredCourses = append(filteredCourses, course)
		} else {
			found := false
			for _, metadataValue := range courseMetadataValues {
				if *metadataValue == *metadataTypeValue {
					found = true
				}
			}
			if !found {
				filteredCourses = append(filteredCourses, course)
			}
		}
	}

	return filteredCourses
}
