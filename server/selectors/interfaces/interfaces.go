package interfaces

import "github.com/eyalkenig/suchef-bot/server/models"

type ICourseSelector interface {
	GetByTheme(theme models.Theme, maxItems int) ([]models.Course, error)
}

type IFilter interface {
	Filter([]*models.Course) []*models.Course
}

type IFilterFactory interface {
	GetFiltersByUserPreference(*models.Preference) []IFilter
}
