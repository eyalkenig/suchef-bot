package filters

import (
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/selectors/interfaces"
)

func GetFiltersByUserPreference(preferences *models.Preference) []interfaces.IFilter {
	filters := []interfaces.IFilter{}

	if preferences.Diet.Name != "anything" {
		filters = append(filters, NewDietFilter(preferences.Diet))
	}

	if preferences.Sensitivity.Name != "no" {
		filters = append(filters, NewSensitivityFilter(preferences.Sensitivity))
	}

	return filters
}
