package theme

import (
	"errors"

	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/repositories"
	"github.com/eyalkenig/suchef-bot/server/selectors"
)

const maxItemsToSelect = 10

func GetSelectedThemeQuickReplies(repository repositories.ICourseRepository,
	selectedTheme *models.Theme,
	userContext context.IUserContext) (map[string]string, error) {
	courses, err := selectors.GetByTheme(repository, userContext, selectedTheme, maxItemsToSelect)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		//TODO: what does this means bazel??
		return nil, errors.New("Not found courses")
	}
	quickReplies := make(map[string]string)

	for _, course := range courses {
		quickReplies[course.Name] = course.ImageURL
	}
	return quickReplies, nil
}
