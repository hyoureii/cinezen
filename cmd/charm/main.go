package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	choices []string
	cursor  int
}

func initialModel() model {
	return model{
		choices: []string{"Cari Film", "Tambah Film Baru", "Tampilkan Daftar Film"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {

}
