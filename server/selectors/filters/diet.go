package filters

import (
	"github.com/eyalkenig/suchef-bot/server/models"
)

type DietFilter struct {
	diet *models.Diet
}

func NewDietFilter(diet *models.Diet) *DietFilter {
	return &DietFilter{diet: diet}
}

func (dietFilter *DietFilter) Filter(courses []*models.Course) []*models.Course {
	return SelectMetadata(courses, models.DietMetadataTypeName, &dietFilter.diet.Name)
}
