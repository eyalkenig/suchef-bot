package filters

import (
	"github.com/eyalkenig/suchef-bot/server/models"
)

type SensitivityFilter struct {
	sensitivity *models.Sensitivity
}

func NewSensitivityFilter(sensitivity *models.Sensitivity) *SensitivityFilter {
	return &SensitivityFilter{sensitivity: sensitivity}
}

func (sensitivityFilter *SensitivityFilter) Filter(courses []*models.Course) []*models.Course {
	return ExcludeMetadata(courses, models.SensitivityMetadataTypeName, &sensitivityFilter.sensitivity.Name)
}
