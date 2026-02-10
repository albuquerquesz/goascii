package tui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"golang.design/x/clipboard"
	"golang.org/x/term"
)

func PromptActions(asciiOutput string) {
	fmt.Fprintln(os.Stdout, "")
	fmt.Fprintln(os.Stdout, "Actions: [C] Copy  [O] Save to file  [Enter] Exit")

	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		return
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return
	}
	defer func() {
		_ = term.Restore(fd, oldState)
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return
		}

		switch b {
		case '\r', '\n':
			return
		case 'c', 'C':
			copyToClipboard(asciiOutput)
		case 'o', 'O':
			path := promptPath()
			if strings.TrimSpace(path) == "" {
				fmt.Fprintln(os.Stderr, "No output path provided")
				continue
			}
			if err := os.WriteFile(path, []byte(asciiOutput), 0o644); err != nil {
				fmt.Fprintln(os.Stderr, "Failed to write file:", err)
			} else {
				fmt.Fprintln(os.Stdout, "Saved to", path)
			}
		}
	}
}

func copyToClipboard(text string) {
	if err := clipboard.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Clipboard not available:", err)
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
	fmt.Fprintln(os.Stdout, "Copied to clipboard")
}

func promptPath() string {
	var path string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Output file path").Value(&path),
		),
	)
	_ = form.Run()
	return path
}
