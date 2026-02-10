package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/albuquerquesz/goascii/internal/ascii"
	"github.com/albuquerquesz/goascii/internal/tui"
	"github.com/charmbracelet/huh"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <text>",
	Short: "Create ASCII art from text",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := strings.Join(args, " ")

		font, _ := cmd.Flags().GetString("font")
		outPath, _ := cmd.Flags().GetString("out")
		isTTY := isatty.IsTerminal(os.Stdin.Fd())

		if font == "" {
			if !isTTY {
				return errors.New("non-interactive mode requires --font")
			}
			selected, err := promptFont()
			if err != nil {
				return err
			}
			font = selected
		}

		if !ascii.IsValidFont(font) {
			return fmt.Errorf("invalid font: %s", font)
		}

		output := ascii.Render(text, font)
		fmt.Fprint(os.Stdout, output)

		if outPath != "" {
			if err := os.WriteFile(outPath, []byte(output), 0o644); err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}
		}

		if isTTY {
			tui.PromptActions(output)
		}

		return nil
	},
}

func init() {
	createCmd.Flags().StringP("font", "f", "", "Font name")
	createCmd.Flags().StringP("out", "o", "", "Write output to a file")
}

func promptFont() (string, error) {
	options := make([]huh.Option[string], 0, len(ascii.Fonts))
	for _, f := range ascii.Fonts {
		options = append(options, huh.NewOption(f, f))
	}

	selected := ascii.DefaultFont
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a font").
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return "", err
	}
	return selected, nil
}
