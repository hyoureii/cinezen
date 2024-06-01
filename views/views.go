package views

import (
	"cinezen/db"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// sementara pake variable global, ntar diganti jadi fungsi
var (
	width     = 96
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	titleStyle = lipgloss.NewStyle().
			Width(50).Align(lipgloss.Center).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(special).
			Padding(1, 0).
			Italic(true).
			Bold(true)

	baseStyle = lipgloss.NewStyle().
			Padding(0, 1)
	headerStyle = baseStyle.
			Foreground(lipgloss.Color("#9ec047")).
			Bold(true)
	dimStyle = baseStyle.
			Foreground(lipgloss.Color("252"))
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Mengubah format angka ke nama bulan untuk print nama bulan
func convertMonth(data db.MovieDB, index int) string {
	switch data.Db[index].Schedule.Month {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	}
	return ""
}

// Menampilkan daftar seluruh film
func ListMovie(data db.MovieDB) {
	fmt.Println(lipgloss.PlaceHorizontal(width, lipgloss.Center, baseStyle.Foreground(special).Italic(true).Render("Database Film")))
	headers := []string{"#", "Judul", "Durasi", "Genre", "Rating", "Jadwal Tayang", "Harga"}

	t := table.New().
		Width(width).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(highlight)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return headerStyle
			case row%2 == 0:
				return baseStyle
			default:
				return dimStyle
			}
		}).
		Headers(headers...)

	for i := range data.Db {
		t.Row(fmt.Sprintf("%d", i+1), data.Db[i].Title, fmt.Sprintf("%d", data.Db[i].Duration), data.Db[i].Genre, fmt.Sprintf("%.1f", data.Db[i].Rating), fmt.Sprintf("%d:00 %d %s", data.Db[i].Schedule.Hour, data.Db[i].Schedule.Day, convertMonth(data, i)), fmt.Sprintf("Rp. %d", data.Db[i].Price))
	}
	fmt.Println(t)
}

func renderTitle(str string) {
	fmt.Println(lipgloss.Place(width, 7,
		lipgloss.Center, lipgloss.Center,
		titleStyle.Render(str),
		lipgloss.WithWhitespaceChars("シネゼン"),
		lipgloss.WithWhitespaceForeground(subtle),
	))
}

func ViewStart() {
	clearScreen()
	renderTitle("Welcome to CINEZEN")
	fmt.Println(baseStyle.Render("\n1. Masuk sebagai Admin\n2. Masuk sebagai User"))
}

func ViewAdmin(showList bool, data db.MovieDB) {
	clearScreen()
	renderTitle("CINEZEN - Admin")
	if showList {
		fmt.Print("\n")
		ListMovie(data)
		fmt.Println(baseStyle.Render("\n1. Sembunyikan Daftar Film"))
	} else {
		fmt.Println(baseStyle.Render("\n1. Tampilkan Daftar Film"))
	}
	fmt.Println(baseStyle.Render("2. Cari film\n3. Edit data film\n4. Tambah film baru"))
}

func ViewUser(showList bool, data db.MovieDB) {
	clearScreen()
	renderTitle("CINEZEN")
	if showList {
		fmt.Print("\n")
		ListMovie(data)
		fmt.Println(baseStyle.Render("\n1. Sembunyikan Daftar Film"))
	} else {
		fmt.Println(baseStyle.Render("\n1. Tampilkan Daftar Film"))
	}
	fmt.Println(baseStyle.Render("2. Cari film\n3. Pesan tiket"))
}

func ViewSearch(data db.MovieDB) {
	clearScreen()
	var opt string
	cursor := 0
	var selectSearch [3]string
	for i := range selectSearch {
		switch i {
		case 0:
			opt = "Judul"
		case 1:
			opt = "Genre"
		case 2:
			opt = "Jadwal Tayang"
		}
		if i == cursor {
			selectSearch[i] = fmt.Sprint(baseStyle.Padding(0, 1).Background(highlight).Render(opt))
		} else {
			selectSearch[i] = fmt.Sprint(dimStyle.Padding(0, 1).Render(opt))
		}
	}
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center, selectSearch[0], selectSearch[1], selectSearch[2]))
	// var dbFound db.MovieDB

}
