package selectors

import (
	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/repositories"
	"github.com/eyalkenig/suchef-bot/server/selectors/filters"
	"github.com/eyalkenig/suchef-bot/server/selectors/interfaces"
	"github.com/eyalkenig/suchef-bot/server/selectors/sorters"
	"math"
)

func GetByTheme(courseRepository repositories.ICourseRepository, userContext context.IUserContext, theme *models.Theme, maxItems int) ([]*models.Course, error) {
	preferences, err := userContext.GetPreferences()
	if err != nil {
		return nil, err
	}

	courseFilters := filters.GetFiltersByUserPreference(preferences)
	courseFilters = append(courseFilters, filters.NewThemeFilter(theme))

	allCourses := courseRepository.GetCourses()

	filteredCourses := filter(allCourses, courseFilters)
	sortedCourses := sorters.SortByMostPopular(filteredCourses)

	numOfItemsToTake := math.Min(float64(maxItems), float64(len(sortedCourses)))
	return sortedCourses[:int(numOfItemsToTake)], nil
}

func filter(courses []*models.Course, filters []interfaces.IFilter) []*models.Course {
	filteredCourses := courses
	for _, filter := range filters {
		filteredCourses = filter.Filter(filteredCourses)
	}
	return filteredCourses
}
