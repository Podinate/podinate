// tui contains the code for the terminal user interface.
package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MODEL DATA

type simplePage struct{ text string }

func NewSimplePage(text string) simplePage {
	return simplePage{text: text}
}

func (s simplePage) Init() tea.Cmd { return nil }

// VIEW

func (s simplePage) View() string {
	textLen := len(s.text)
	topAndBottomBar := strings.Repeat("*", textLen+4)
	return fmt.Sprintf(
		"%s\n* %s *\n%s\n\nPress Ctrl+C to exit",
		topAndBottomBar, s.text, topAndBottomBar,
	)
}

// UPDATE

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return s, tea.Quit
		}
	}
	return s, nil
}

var (
	StyleSuccess = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#66ff33")).
			Padding(0, 1)
	StyleItalic = lipgloss.NewStyle().
			Italic(true)
	StyleUpdated = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#ffcc00")).
			Padding(0, 1)
	StyleError = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#d00000")).
			Padding(0, 1)
)
