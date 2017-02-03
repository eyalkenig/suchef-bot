package models

type Course struct {
	ID          int64
	Name        string
	ImageURL    string
	Description string

	Ingredients []*Ingredient

	Tags map[string][]*string
}

func NewCourse(id int64, name, description, imageURL string) *Course {
	course := &Course{ID: id, Name: name, ImageURL: imageURL, Description: description}
	course.Tags = make(map[string][]*string)
	return course
}
