package views

import (
	"cinezen/db"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func getStyle() lipgloss.Style {
	return lipgloss.NewStyle().Padding(0, 1)
}

func getColor(opt string) lipgloss.Color {
	switch opt {
	case "green":
		return lipgloss.Color("#24ff86")
	case "purple":
		return lipgloss.Color("#874BFD")
	case "cyan":
		return lipgloss.Color("#6fedff")
	case "gray":
		return lipgloss.Color("245")
	case "teal":
		return lipgloss.Color("#00acc1")
	case "red":
		return lipgloss.Color("#db4b4b")
	case "black":
		return lipgloss.Color("#030303")
	}
	return ""
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Menampilkan daftar seluruh film
func ListMovie(data db.MovieDB, str string) {
	width := 115
	style := getStyle()
	header := style.Bold(true).Foreground(getColor("green"))
	fmt.Println(lipgloss.PlaceHorizontal(width, lipgloss.Center, header.Italic(true).Render(str)))
	headers := [8]string{"#", "Judul", "Durasi", "Genre", "Rating", "Jadwal Tayang", "Harga", "Discount"}

	t := table.New().
		Width(width).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(getColor("teal"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return header
			case row%2 == 0:
				return style
			default:
				return style.Foreground(lipgloss.Color("252"))
			}
		}).
		Headers(headers[0], headers[1], headers[2], headers[3], headers[4], headers[5], headers[6], headers[7])

	for i := 0; i < data.Len; i++ {
		t.Row(fmt.Sprintf("%d", i+1), data.Db[i].Title, fmt.Sprintf("%d Menit", data.Db[i].Duration), data.Db[i].Genre, fmt.Sprintf("%.1f", data.Db[i].Rating), fmt.Sprintf("%d:00 %d %s", data.Db[i].Schedule.Hour, data.Db[i].Schedule.Date, db.ConvertMonth(data.Db[i].Schedule.Month)), fmt.Sprintf("Rp. %d", data.Db[i].Price), fmt.Sprintf("%d%%", data.Db[i].Discount))
	}
	fmt.Println(t)
}

func RenderTitle(str string, width, padY, marginY int) {
	screenWidth := 115
	style := lipgloss.NewStyle().
		Width(width).Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(getColor("green")).
		Padding(padY, 0).Italic(true)
	fmt.Println(lipgloss.Place(screenWidth, ((padY+marginY)*2)+3,
		lipgloss.Center, lipgloss.Center,
		style.Render(str),
		lipgloss.WithWhitespaceChars("シネゼン"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#4c2f94")),
	))
	fmt.Println(lipgloss.NewStyle().
		Width(screenWidth).
		BorderForeground(getColor("teal")).
		BorderTop(true).
		BorderStyle(lipgloss.Border{Top: "-"}))
}

func RenderTip(str string) {
	style := lipgloss.NewStyle().Italic(true).Foreground(getColor("gray"))
	fmt.Println(style.Render(str))
}

func PrintError(str string) {
	style := lipgloss.NewStyle().Foreground(getColor("red"))
	fmt.Println(style.Render("\nERROR : " + str))
}

func ViewStart() {
	style := getStyle()
	clearScreen()
	RenderTitle("Welcome to CINEZEN", 50, 1, 1)
	fmt.Println(style.Render("\n1. Masuk sebagai Admin\n2. Masuk sebagai User"))
	RenderTip("\nTekan q untuk keluar dari aplikasi")
}

func ViewAdmin(showList bool, data db.MovieDB) {
	style := getStyle()
	clearScreen()
	RenderTitle("CINEZEN - Admin", 50, 1, 1)
	if showList {
		fmt.Println()
		ListMovie(data, "Database Film")
		fmt.Println(style.Render("\n1. Sembunyikan Daftar Film"))
	} else {
		fmt.Println(style.Render("\n1. Tampilkan Daftar Film"))
	}
	fmt.Println(style.Render("2. Cari film\n3. Edit data film\n4. Tambah film baru\n5. Hapus data film"))
	if showList {
		fmt.Println(style.Render("\n7. Urutkan Berdasarkan Genre\n8. Urutkan Berdasarkan Jadwal Tayang\n9. Urutkan Berdasarkan Harga"))
	}

	RenderTip("\nTekan q untuk log out")
}

func ListTicket(data db.Tickets, str string) {
	width := 115
	style := getStyle()
	header := style.Bold(true).Foreground(getColor("green"))
	fmt.Println(lipgloss.PlaceHorizontal(width, lipgloss.Center, header.Italic(true).Render(str)))
	headers := [8]string{"#", "Judul", "Durasi", "Genre", "Rating", "Jadwal Tayang", "Harga", "Discount"}

	t := table.New().
		Width(width).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(getColor("teal"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return header
			case row%2 == 0:
				return style
			default:
				return style.Foreground(lipgloss.Color("252"))
			}
		}).
		Headers(headers[0], headers[1], headers[2], headers[3], headers[4], headers[5], headers[6], headers[7])

	for i := 0; i < data.Movies.Len; i++ {
		t.Row(fmt.Sprintf("%d", i+1), data.Movies.Db[i].Title, fmt.Sprintf("%d Menit", data.Movies.Db[i].Duration), data.Movies.Db[i].Genre, fmt.Sprintf("%.1f", data.Movies.Db[i].Rating), fmt.Sprintf("%d:00 %d %s", data.Movies.Db[i].Schedule.Hour, data.Movies.Db[i].Schedule.Date, db.ConvertMonth(data.Movies.Db[i].Schedule.Month)), fmt.Sprintf("Rp. %d", data.Movies.Db[i].Price), fmt.Sprintf("%d%%", data.Movies.Db[i].Discount))
	}
	fmt.Println(t)
}

func ViewUser(showList, showTickets bool, ticket db.Tickets, data db.MovieDB) {
	style := getStyle()
	clearScreen()
	RenderTitle("CINEZEN", 50, 1, 1)
	if showTickets {
		ListTicket(ticket, "Tiket yang dimiliki")
	}
	if showList {
		fmt.Print("\n")
		ListMovie(data, "Daftar Film")
		fmt.Println(style.Render("\n1. Sembunyikan Daftar Film"))
	} else {
		fmt.Println(style.Render("\n1. Tampilkan Daftar Film"))
	}
	fmt.Println(style.Render("2. Cari film\n3. Pesan tiket"))
	if showTickets {
		fmt.Println(style.Render("4. Sembunyikan Daftar Tiket"))
	} else {
		fmt.Println(style.Render("4. Tampilkan Daftar Tiket"))
	}
	if showList {
		fmt.Println(style.Render("\n7. Urutkan Berdasarkan Genre\n8. Urutkan Berdasarkan Jadwal Tayang\n9. Urutkan Berdasarkan Harga"))
	}

	RenderTip("\nTekan q untuk log out")
}

func ViewAdd(phase int, data db.Movies, success bool) {
	input := getStyle().Padding(0, 0).Foreground((getColor("cyan")))
	clearScreen()
	RenderTitle("Tambah Film", 45, 0, 0)
	if success {
		var dataF db.MovieDB
		dataF.Db[0] = data
		dataF.Len = 1
		ListMovie(dataF, "Film Ditambahkan")
		fmt.Println("Film berhasil ditambahkan!")
		RenderTip("\nENTER untuk menambah film lagi, ESC untuk kembali ke menu utama")
	} else {
		if data.Title == "" {
			RenderTip("\ngunakan _ (underscore) untuk spasi")
		}
		fmt.Printf("Judul : %s", input.Render(data.Title))
		if phase > 0 {
			if data.Duration != 0 {
				fmt.Printf("\nDurasi : %s\n", input.Render(fmt.Sprintf("%d", data.Duration)))
			} else {
				RenderTip("\n\ndurasi dalam menit")
				fmt.Print("Durasi : ")
			}
		}
		if phase > 1 {
			genres := db.GetGenres()
			if data.Genre == "" {
				fmt.Println("\nPilih genre :")
				for i := range genres {
					if i == 9 {
						genres[i] = getStyle().Padding(0, 2).Render(fmt.Sprintf("0. %s", genres[i]))
					} else {
						genres[i] = getStyle().Padding(0, 2).Render(fmt.Sprintf("%d. %s", i+1, genres[i]))
					}
				}
				fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.JoinVertical(lipgloss.Left, genres[0], genres[1], genres[2], genres[3], genres[4]), lipgloss.JoinVertical(lipgloss.Left, genres[5], genres[6], genres[7], genres[8], genres[9])))
				RenderTip("\npilih menggunakan angka")
			}
			fmt.Printf("Genre : %s", input.Render(data.Genre))
		}
		if phase > 2 {
			if data.Rating != 0 {
				fmt.Printf("\nRating : %s\n", input.Render(fmt.Sprintf("%.1f", data.Rating)))
			} else {
				RenderTip("\n\nrating 1 - 10")
				fmt.Print("Rating : ")
			}
		}
		if phase > 3 {
			if data.Price != 0 {
				fmt.Printf("Harga :  %s\n", input.Render(fmt.Sprintf("%d", data.Price)))
			} else {
				RenderTip("\nmin. harga 30.000")
				fmt.Print("Harga : ")
			}
		}
		if phase > 4 {
			if data.Schedule.Month == 0 {
				RenderTip("\ngunakan angka, contoh 2 untuk Februari")
			}
			fmt.Printf("Bulan tayang : %s", input.Render(db.ConvertMonth(data.Schedule.Month)))
		}
		if phase > 5 {
			if data.Schedule.Date != 0 {
				fmt.Printf("\nTanggal tayang : %s\n", input.Render(fmt.Sprintf("%d", data.Schedule.Date)))
			} else {
				RenderTip("\n\nmasukkan tanggal 1 - 31")
				fmt.Printf("Tanggal tayang : ")
			}
		}
		if phase > 6 {
			if data.Schedule.Hour != 0 {
				fmt.Printf("Jam Tayang : %s\n", input.Render(fmt.Sprintf("%d.00", data.Schedule.Hour)))
			} else {
				RenderTip("\njam tayang hanya 10.00 - 21.00")
				fmt.Printf("Jam Tayang : ")
			}
		}
		if phase > 7 {
			var dataP db.MovieDB
			dataP.Db[0] = data
			dataP.Len = 1
			ListMovie(dataP, "")
			fmt.Println("Yakin ingin menambah film ?")
			RenderTip("ENTER untuk konfirmasi, TAB untuk ulang input data, ESC untuk cancel penambahan film\n")
		}
	}
}

func ViewSearch(data db.MovieDB, mode int, str string) {
	style := getStyle()
	dim := style.Foreground((getColor("gray"))).Italic(true)
	highlight := style.Background((getColor("green"))).Foreground(getColor("black"))
	input := style.Foreground((getColor("cyan")))
	clearScreen()
	RenderTitle("Cari Film", 40, 0, 0)
	ListMovie(data, "")

	fmt.Print(dim.Render("\nmenampilkan hasil :"))
	fmt.Printf("\"%s\"", input.Render(str))
	fmt.Print("\nCari berdasarkan ")
	switch mode {
	case 0:
		fmt.Print(highlight.Render("Judul"))
	case 1:
		fmt.Print(highlight.Render("Genre"))
	case 2:
		fmt.Print(highlight.Render("Jadwal Tayang"))
	}
	fmt.Print(dim.Render("(TAB untuk mengubah mode pencarian)"))

	RenderTip("\n\nESC untuk kembali ke menu utama\n")
	switch mode {
	case 0:
		fmt.Printf("Cari : %s", input.Render(str))
	case 1:
		genres := db.GetGenres()
		for i := range genres {
			if i == 9 {
				genres[i] = getStyle().Padding(0, 2).Render(fmt.Sprintf("0. %s", genres[i]))
			} else {
				genres[i] = getStyle().Padding(0, 2).Render(fmt.Sprintf("%d. %s", i+1, genres[i]))
			}
		}
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.JoinVertical(lipgloss.Left, genres[0], genres[1], genres[2], genres[3], genres[4]), lipgloss.JoinVertical(lipgloss.Left, genres[5], genres[6], genres[7], genres[8], genres[9])))
		RenderTip("\nPilih genre menggunakan angka")
		fmt.Printf("\nCari : %s", input.Render(str))
	case 2:
		RenderTip("Masukkan bulan menggunakan angka, contoh \"19 2 2\" untuk 19.00 2 Februari")
		fmt.Printf("\nCari : %s", input.Render(str))
	}
}

func ViewEdit(data db.MovieDB, idChosen, dataChosen bool, id, dataType int) {
	clearScreen()
	RenderTitle("Edit Data Film", 45, 0, 0)
	style := getStyle()
	highlight := style.Foreground(getColor("cyan"))
	var chosen db.Movies
	if id > -1 {
		chosen = data.Db[id]
	} else {
		chosen = db.Movies{Title: "", Duration: 0, Genre: "", Rating: 0, Price: 0, Schedule: db.MovieSchedule{Hour: 0, Date: 0, Month: 0}}
	}
	if !idChosen {
		ListMovie(data, "Database Film")
		RenderTip("\nPilih q untuk kembali ke menu utama")
		fmt.Print("\nPilih no. urut film yang akan diedit : ")
	} else {
		var dbChosen db.MovieDB
		dbChosen.Db[0] = chosen
		dbChosen.Len = 1
		ListMovie(dbChosen, "")
		if !dataChosen {
			fmt.Println(style.Render("\n1. Judul"))
			fmt.Println(style.Render("2. Durasi"))
			fmt.Println(style.Render("3. Genre"))
			fmt.Println(style.Render("4. Rating"))
			fmt.Println(style.Render("5. Harga"))
			fmt.Println(style.Render("6. Jadwal Tayang"))
			RenderTip("\nPilih q untuk kembali")
			fmt.Print("\nPilih data yang mau diubah : ")
		} else {
			last := data.Db[id]
			switch dataType {
			case 0:
				fmt.Printf("Judul lama : %s\n", last.Title)
				fmt.Print("\nJudul baru : ")
			case 1:
				fmt.Printf("Durasi lama : %d\n", last.Duration)
				fmt.Print("\nDurasi baru : ")
			case 2:
				fmt.Printf("Genre lama : %s\n", last.Genre)
				fmt.Println(style.Render("\n1. Action"))
				fmt.Println(style.Render("\n2. Comedy"))
				fmt.Println(style.Render("\n3. Drama"))
				fmt.Println(style.Render("\n4. Horror"))
				fmt.Println(style.Render("\n5. Romance"))
				fmt.Println(style.Render("\n6. Sci-Fi"))
				fmt.Println(style.Render("\n7. Documentary"))
				fmt.Println(style.Render("\n8. Thriller"))
				fmt.Println(style.Render("\n9. Animation"))
				fmt.Println(style.Render("\n0. Fantasy"))
				RenderTip("\npilih genre menggunakan angka")
				fmt.Print("Genre baru : ")
			case 3:
				fmt.Printf("Rating lama : %.2f\n", last.Rating)
				fmt.Print("\nRating baru : ")
			case 4:
				fmt.Printf("Harga lama : %d\n", last.Price)
				fmt.Print("\nHarga baru : ")
			case 5:
				fmt.Printf("Jadwal tayang lama : %d.00 %d %s\n", last.Schedule.Hour, last.Schedule.Date, db.ConvertMonth(last.Schedule.Month))
				fmt.Println("\nJadwal tayang baru :")
				fmt.Print(style.Render("Bulan : "))
			case 6:
				fmt.Printf("Jadwal tayang lama : %d.00 %d %s\n", last.Schedule.Hour, last.Schedule.Date, db.ConvertMonth(last.Schedule.Month))
				fmt.Println("\nJadwal tayang baru :")
				fmt.Print(style.Render(fmt.Sprintf("Bulan : %s", highlight.Render(db.ConvertMonth(data.Db[id].Schedule.Month)))))
				fmt.Print(style.Render("\nTanggal : "))
			case 7:
				fmt.Printf("Jadwal tayang lama : %d.00 %d %s\n", last.Schedule.Hour, last.Schedule.Date, db.ConvertMonth(last.Schedule.Month))
				fmt.Println("\nJadwal tayang baru :")
				fmt.Print(style.Render(fmt.Sprintf("Bulan : %s", highlight.Render(db.ConvertMonth(data.Db[id].Schedule.Month)))))
				fmt.Print(style.Render(fmt.Sprintf("\nTanggal : %s", highlight.Render(fmt.Sprintf("%d", data.Db[id].Schedule.Date)))))
				fmt.Print(style.Render("\nJam : "))
			}
		}
	}
}

func ViewDelete(data db.MovieDB, id int, hasChosen bool) {
	clearScreen()
	var chosen db.Movies
	if id > -1 {
		chosen = data.Db[id]
	} else {
		chosen = db.Movies{Title: "", Duration: 0, Genre: "", Rating: 0, Price: 0, Discount: 0, Schedule: db.MovieSchedule{Hour: 0, Date: 0, Month: 0}}
	}
	if hasChosen {
		var dbChosen db.MovieDB
		dbChosen.Db[0] = chosen
		dbChosen.Len = 1
		ListMovie(dbChosen, "")
		fmt.Println("\nYakin ingin menhapus data film ini ?")
		RenderTip("ENTER untuk konfirmasi, ESC untuk cancel")
	} else {
		ListMovie(data, "Database Film")
		RenderTip("\nPilih 0 untuk kembali ke menu utama")
		fmt.Print("\nPilih data film yang ingin dihapus : ")
	}
}

func BuyTicket(data db.MovieDB, id int, idChosen, success bool) {
	clearScreen()
	var chosen db.Movies
	if id > -1 {
		chosen = data.Db[id]
		chosen.Price = chosen.Price - ((chosen.Price * chosen.Discount) / 100)
	} else {
		chosen = db.Movies{Title: "", Duration: 0, Genre: "", Rating: 0, Price: 0, Discount: 0, Schedule: db.MovieSchedule{Hour: 0, Date: 0, Month: 0}}
	}
	var dbChosen db.MovieDB
	dbChosen.Db[0] = chosen
	dbChosen.Len = 1
	if success {
		RenderTitle("Film Berhasil Dibeli!", 45, 0, 0)
		ListMovie(dbChosen, "")
		fmt.Println("\nApakah anda ingin membeli tiket lain?")
		RenderTip("ENTER untuk konfirmasi, ESC untuk cancel")
	} else {
		if !idChosen {
			ListMovie(data, "")
			RenderTip("\nPilih 0 untuk kembali ke menu utama")
			fmt.Print("Pilih film yang ingin dibeli : ")
		} else {
			ListMovie(dbChosen, "")
			fmt.Println("\nApakah anda ingin membeli tiket untuk film ini?")
			RenderTip("ENTER untuk konfirmasi, ESC untuk cancel")
		}
	}
}

func SeatSelect(seat db.AvailSeat) {
	clearScreen()
	width := 115
	filled, empty := "▨", "▢"
	var seats [20]string
	t := table.New().
		Width(width).
		Border(lipgloss.HiddenBorder()).
		Headers(" ", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	for i := range seat {
		for j := range seat[i] {
			if !seat[i][j] {
				seats[j] = empty
			} else {
				seats[j] = getStyle().Padding(0, 0).Foreground(getColor("teal")).Render(filled)
			}
		}
		t.Row(fmt.Sprintf("%c", byte(i+65)), seats[0], seats[1], seats[2], seats[3], seats[4], seats[5], seats[6], seats[7], seats[8], seats[9], seats[10], seats[11], seats[12], seats[13], seats[14], seats[15], seats[16], seats[17], seats[18], seats[19])
	}
	fmt.Println(t)
	RenderTip("\nBiru berarti terisi")
	fmt.Println("Silahkan pilih tempat duduk : ")
}
