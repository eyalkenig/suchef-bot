package models

type IModelFactory interface {
	GetDiet(dietID int64) (*Diet, error)
	GetSensitivity(sensitivityID int64) (*Sensitivity, error)
}
