package inputs

type ThemeAsian struct{}

func NewThemeAsian() *ThemeAsian {
	return &ThemeAsian{}
}

const THEME_ASIAN_INPUT = "assian-theme"
const THEME_ASIAN_TITLE = "אסייתי 🍜"

func (theme *ThemeAsian) Payload() string {
	return THEME_ASIAN_INPUT
}
