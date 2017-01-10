package inputs

type ThemeMoroccan struct{}

func NewThemeMoroccan() *ThemeMoroccan {
	return &ThemeMoroccan{}
}

const THEME_MOROCCAN_INPUT = "moroccan-theme"
const THEME_MOROCCAN_TITLE = "מרוקאי 🌋"

func (theme *ThemeMoroccan) Payload() string {
	return THEME_MOROCCAN_INPUT
}
