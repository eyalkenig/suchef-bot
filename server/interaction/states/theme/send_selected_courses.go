package theme

import (
	"errors"

	"github.com/eyalkenig/suchef-bot/server/interaction/context"
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/repositories"
	"github.com/eyalkenig/suchef-bot/server/selectors"

	"github.com/eyalkenig/suchef-bot/server/interfaces/messaging"
	conceteCards "github.com/eyalkenig/suchef-bot/server/interaction/cards"
)

const maxItemsToSelect = 10

func GetSelectedThemeCourses(repository repositories.ICourseRepository,
	selectedTheme *models.Theme,
	userContext context.IUserContext) ([]messaging.ICard, error) {
	courses, err := selectors.GetByTheme(repository, userContext, selectedTheme, maxItemsToSelect)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		//TODO: what does this means bazel??
		return nil, errors.New("Not found courses")
	}

	var cards []messaging.ICard

	for _, course := range courses {
		courseCard := conceteCards.NewCourse(course.ID, course.Name, course.Description, course.ImageURL)
		cards = append(cards, courseCard)
	}
	return cards, nil
}

func GetAlterantivesCourses() map[string]string {
	others := make(map[string]string)

	return others

	others["משהו אחר"] = "i_want_something_else"
	others["!תבחרי את"] = "pick_something_for_me"

	return others
}