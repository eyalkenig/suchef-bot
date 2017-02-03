package cards

import (
	"github.com/eyalkenig/suchef-bot/server/interfaces/messaging"
	"github.com/eyalkenig/suchef-bot/server/interaction/buttons"
)

type Course struct {
	id int64
	name string
	description string
	imageURL string
}

func NewCourse(id int64, name, description, imageURL string) *Course {
	return &Course{id: id,name: name, description: description, imageURL: imageURL}
}

func (c *Course) Title() string {
	return c.name
}

func (c *Course) Subtitle() string {
	return c.description
}

func (c *Course) ImageURL() string {
	return c.imageURL
}

func (c *Course) Buttons() []messaging.IButton {
	return []messaging.IButton{buttons.NewIngredients(c.id)}
}

