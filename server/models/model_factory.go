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

func GetSensitivity(sensitivityID int64) (*Sensitivity, error) {
	var sensitivity *Sensitivity
	switch sensitivityID {
	case 0:
		sensitivity = &Sensitivity{ID: 0, Name: "no"}
	case 10:
		sensitivity = &Sensitivity{ID: 10, Name: "gluten"}
	case 20:
		sensitivity = &Sensitivity{ID: 20, Name: "milk"}
	default:
		return nil, errors.New("no such sensitivity id")
	}
	return sensitivity, nil
}
