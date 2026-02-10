package ascii

import "testing"

func TestRenderNotEmpty(t *testing.T) {
	out := Render("test", DefaultFont)
	if out == "" {
		t.Fatal("expected output")
	}
}
