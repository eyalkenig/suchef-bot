package repositories

import (
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/providers"
	"fmt"
)

type CourseRepository struct {
	courses []*models.Course
}

func NewCourseRepository(dataProvider providers.IBotDataProvider) *CourseRepository {
	courses, err := dataProvider.FetchCourses()

	if err != nil {
		fmt.Println("Failed to load courses.")
		courses = []*models.Course{}
	}

	fmt.Println(fmt.Sprintf("Successfully loaded courses. number of courses: %d", len(courses)))
	return &CourseRepository{courses: courses}
}

func (repository *CourseRepository) GetCourses() []*models.Course {
	return repository.courses
}
