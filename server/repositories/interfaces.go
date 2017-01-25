package repositories

import "github.com/eyalkenig/suchef-bot/server/models"

type ICourseRepository interface {
	GetCourses() []*models.Course
}
