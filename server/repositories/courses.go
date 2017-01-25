package repositories

import (
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/providers"
)

type CourseRepository struct {
	courses []*models.Course
}

func NewCourseRepository(dataProvider providers.IBotDataProvider) *CourseRepository {
	courses, err := dataProvider.FetchCourses()

	if err != nil {
		courses = []*models.Course{}
	}

	return &CourseRepository{courses: courses}
}

func (repository *CourseRepository) GetCourses() []*models.Course {
	return repository.courses
}
