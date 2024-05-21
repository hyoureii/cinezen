package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Choice struct {
	Title   string
	Msg     string
	Choices []string
}

var Choose Choice

// storing variable for styling
var (
	defaultStyle = lipgloss.NewStyle()
	titleStyle   = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			Padding(0, 2, 0, 2).
			BorderForeground(lipgloss.Color("#ffffff"))
	selectStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#65141a"))
)

func (c *Choice) ViewStart(cursor int) string {
	c.Title = titleStyle.Render("Welcome to CINEZEN")
	c.Msg = "Anda belum login.\nLogin dulu lah anying."
	c.Choices = []string{
		"Masuk sebagai Pengguna",
		"Masuk sebagai Admin",
	}

	return c.renderChoices(cursor)
}

func (c Choice) renderChoices(cursor int) string {
	s := fmt.Sprintf("%s\n\n%s\n\n", c.Title, c.Msg)
	for i, idChoice := range c.Choices {
		cursorDisp := "   "
		if cursor == i {
			cursorDisp = " ▶ "
			s += fmt.Sprintf("%s%s\n", cursorDisp, selectStyle.Render(idChoice))
		} else {
			s += fmt.Sprintf("%s%s\n", cursorDisp, defaultStyle.Render(idChoice))
		}
	}

	return s + lipgloss.NewStyle().Foreground(lipgloss.Color("#686c98")).Render("\n\ntekan esc untuk keluar")
}
