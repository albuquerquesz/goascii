package cmd

import "testing"

func TestCreateNonTTYRequiresFont(t *testing.T) {
	_, err := resolveFont("", false)
	if err == nil {
		t.Fatal("expected error for missing font in non-interactive mode")
	}

	if err.Error() != "non-interactive mode requires --font" {
		t.Fatalf("unexpected error: %v", err)
	}
}
