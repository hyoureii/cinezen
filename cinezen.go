package main

import (
	"cinezen/app"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Cursor int
	Inputs []textinput.Model
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
	return textinput.Blink
}

// main Update function
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" {
			return m, tea.Quit
		}
	}
	if app.WhatToShow == "choices" {
		return updateChoices(msg, m)
	}

	return m, nil
}

// main View function
func (m Model) View() string {
	var s string

	switch app.AppState {
	case "notLogged":
		s = app.Choose.ViewStart(m.Cursor)
	case "mainAdmin":
		s = app.Choose.ViewAdmin(m.Cursor)
	}

	return s + fmt.Sprintf("\ncursor=%d\nAppState=%s", m.Cursor, app.AppState)
}

func (m *Model) ResetCursor() {
	m.Cursor = 0
}

// Sub-update functions

func updateChoices(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
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
			m.ResetCursor()
		}
	}

	return m, nil
}
