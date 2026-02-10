package ascii

import "testing"

func TestDefaultFontIsValid(t *testing.T) {
	if !IsValidFont(DefaultFont) {
		t.Fatalf("default font %q is not in list", DefaultFont)
	}
}
