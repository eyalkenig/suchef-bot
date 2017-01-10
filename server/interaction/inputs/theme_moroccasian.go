package inputs

type ThemeMoroccasian struct{}

func NewThemeMoroccasian() *ThemeMoroccasian {
	return &ThemeMoroccasian{}
}

const THEME_MOROCCASIAN_INPUT = "moroccasian-theme"
const THEME_MOROCCASIAN_TITLE = "שילוב! 🔥"

func (theme *ThemeMoroccasian) Payload() string {
	return THEME_MOROCCASIAN_INPUT
}
