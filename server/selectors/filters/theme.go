package filters

import (
	"github.com/eyalkenig/suchef-bot/server/models"
)

type ThemeFilter struct {
	theme *models.Theme
}

func NewThemeFilter(theme *models.Theme) *ThemeFilter {
	return &ThemeFilter{theme: theme}
}

func (themeFilter *ThemeFilter) Filter(courses []*models.Course) []*models.Course {
	return SelectMetadata(courses, models.ThemeMetadataTypeName, &themeFilter.theme.Name)
}
