package inputs

type FreeTextInput struct {
	Text string
}

func NewFreeTextInput(text string) *FreeTextInput {
	return &FreeTextInput{Text: text}
}

const FREE_TEXT_INPUT = "free-text"

func (diet *FreeTextInput) Payload() string {
	return FREE_TEXT_INPUT
}
