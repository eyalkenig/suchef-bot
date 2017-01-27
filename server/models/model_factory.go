package models

import "errors"

func GetDiet(dietID int64) (*Diet, error) {
	var diet *Diet
	switch dietID {
	case 0:
		diet = &Diet{ID: 0, Name: "anything"}
	case 10:
		diet = &Diet{ID: 10, Name: "vegan"}
	case 20:
		diet = &Diet{ID: 20, Name: "vegetarian"}
	default:
		return nil, errors.New("no such diet id")
	}
	return diet, nil
}

func GetDietByName(dietName *string) (*Diet, error) {
	var diet *Diet
	switch *dietName {
	case "anything":
		diet = &Diet{ID: 0, Name: "anything"}
	case "vegan":
		diet = &Diet{ID: 10, Name: "vegan"}
	case "vegetarian":
		diet = &Diet{ID: 20, Name: "vegetarian"}
	default:
		return nil, errors.New("no such diet name")
	}
	return diet, nil
}

func GetSensitivity(sensitivityID int64) (*Sensitivity, error) {
	var sensitivity *Sensitivity
	switch sensitivityID {
	case 10:
		sensitivity = &Sensitivity{ID: 10, Name: "gluten"}
	case 20:
		sensitivity = &Sensitivity{ID: 20, Name: "milk"}
	default:
		return nil, errors.New("no such sensitivity id")
	}
	return sensitivity, nil
}

func GetSensitivityByName(sensitivityName *string) (*Sensitivity, error) {
	var sensitivity *Sensitivity
	switch *sensitivityName {
	case "gluten":
		sensitivity = &Sensitivity{ID: 10, Name: "gluten"}
	case "milk":
		sensitivity = &Sensitivity{ID: 20, Name: "milk"}
	default:
		return nil, errors.New("no such sensitivity")
	}
	return sensitivity, nil
}

func GetThemeByName(themeName *string) (*Theme, error) {
	var theme *Theme
	switch *themeName {
	case "asian":
		theme = &Theme{ID: 10, Name: "asian"}
	case "moroccan":
		theme = &Theme{ID: 20, Name: "moroccan"}
	case "moroccasian":
		theme = &Theme{ID: 30, Name: "moroccasian"}
	default:
		return nil, errors.New("no such theme")
	}
	return theme, nil
}