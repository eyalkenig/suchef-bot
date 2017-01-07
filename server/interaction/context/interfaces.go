package context

type IUserContext interface {
	GetID() int64
	GetExternalUserID() string
	IsMale() bool

	SetDiet(dietID int64) error
	SetSensitivity(sensitivityID int64) error
}