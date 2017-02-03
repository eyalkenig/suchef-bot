package buttons

import "fmt"

type Ingredients struct {
	id int64
}

const INGREDIENTS_PAYLOAD = "see_ingrediants_of_course_"

func NewIngredients(id int64) *Ingredients{
	return &Ingredients{id: id}
}

func (i *Ingredients) Type() string {
	return "postback"
}
func (i *Ingredients) Title() string {
	return "מרכיבים"
}
func (i *Ingredients) Payload() string {
	return fmt.Sprintf("%s%d", INGREDIENTS_PAYLOAD, i.id)
}
