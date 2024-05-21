package main

import (
	"cinezen/app"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Cursor int
}

func main() {
	var m Model
	app.InitApp()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// main Update function
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" {
			return m, tea.Quit
		}
	}
	if app.AppState == "notLogged" {
		return updateStart(msg, m)
	}

	return m, nil
}

// main View function
func (m Model) View() string {
	var s string

	if app.AppState == "notLogged" {
		s = app.Choose.ViewStart(m.Cursor)
	}

	return s
}

// Sub-update functions

func updateStart(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(app.Choose.Choices)-1 {
				m.Cursor++
			}
		case "enter":
			app.UpdateApp(m.Cursor)
		}
	}

	return m, nil
}
