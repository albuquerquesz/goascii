package ascii

var (
	DefaultFont = "standard"
	Fonts       = []string{
		"standard",
		"slant",
		"big",
		"small",
		"shadow",
		"block",
		"banner",
		"doom",
	}
)

func IsValidFont(font string) bool {
	for _, f := range Fonts {
		if f == font {
			return true
		}
	}
	return false
}
