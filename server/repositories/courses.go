package repositories

import (
	"fmt"

	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"github.com/eyalkenig/suchef-bot/server/models"
)

type CourseRepository struct {
	courses []*models.Course
}

func NewCourseRepository(accountID int64, dataProvider providers.IBotDataProvider) *CourseRepository {
	courses, err := dataProvider.FetchCourses(accountID)

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
