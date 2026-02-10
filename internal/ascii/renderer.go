package ascii

import (
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func Render(text, font string) string {
	clean := strings.TrimSpace(text)
	fig := figure.NewFigure(clean, font, true)
	return fig.String()
}
