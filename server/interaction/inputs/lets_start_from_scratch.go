package inputs

type LetsStartFromScratch struct {}

func NewLetsStartFromScratch() *LetsStartFromScratch{
	return &LetsStartFromScratch{}
}
const LETS_START_FROM_SCRATCH_INPUT = "lets-start-from-scratch-10192838"
const LETS_START_FROM_SCRATCH_TITLE = "נתחיל מחדש"

func (diet *LetsStartFromScratch) Payload() string{
	return LETS_START_FROM_SCRATCH_INPUT
}
