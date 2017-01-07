package inputs

type SensitivityMilk struct {}

func NewSensitivityMilk() *SensitivityMilk{
	return &SensitivityMilk{}
}

const MILK_SENSITIVITY_INPUT = "milk-sensitive"
const MILK_SENSITIVITY_TITLE = "חלב 🍰"

func (diet *SensitivityMilk) Payload() string{
	return MILK_SENSITIVITY_INPUT
}

