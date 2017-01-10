package inputs

type SensitivityGluten struct{}

func NewSensitivityGluten() *SensitivityGluten {
	return &SensitivityGluten{}
}

const GLUTEN_SENSITIVITY_INPUT = "gluten-sensitive"
const GLUTEN_SENSITIVITY_TITLE = "גלוטן 🍞"

func (diet *SensitivityGluten) Payload() string {
	return GLUTEN_SENSITIVITY_INPUT
}
