package inputs

type DietVegan struct{}

func NewDietVegan() *DietVegan {
	return &DietVegan{}
}

const DIET_VEGAN_INPUT = "vegan-diet"
const DIET_VEGAN_TITLE = "טבעוני🍋"

func (diet *DietVegan) Payload() string {
	return DIET_VEGAN_INPUT
}
