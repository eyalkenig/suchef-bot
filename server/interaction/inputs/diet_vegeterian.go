package inputs

type DietVegetarian struct{}

func NewDietVegetarian() *DietVegetarian {
	return &DietVegetarian{}
}

const DIET_VEGETARIAN_INPUT = "vegetarian-diet"
const DIET_VEGETARIAN_TITLE = "צמחוני 🧀"

func (diet *DietVegetarian) Payload() string {
	return DIET_VEGETARIAN_INPUT
}
