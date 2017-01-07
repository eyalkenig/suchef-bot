package inputs

type DietAnything struct {}

func NewDietAnything() *DietAnything{
	return &DietAnything{}
}

const DIET_ANYTHING_INPUT = "no-diet"
const DIET_ANYTHING_TITLE = "אוכל-כל🍗"

func (diet *DietAnything) Payload() string{
	return DIET_ANYTHING_INPUT
}
