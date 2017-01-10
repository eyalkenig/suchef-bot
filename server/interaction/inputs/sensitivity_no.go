package inputs

type SensitivityNo struct{}

func NewSensitivityNo() *SensitivityNo {
	return &SensitivityNo{}
}

const NO_SENSITIVITY_INPUT = "not-sensitive"
const NO_SENSITIVITY_TITLE = "לא 💪"

func (diet *SensitivityNo) Payload() string {
	return NO_SENSITIVITY_INPUT
}
