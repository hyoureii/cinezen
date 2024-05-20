package main

import (
	"fmt"
)

const MAX_MOV int = 100

var db_Len int

type movie struct {
	Title, Genre       string
	Duration, Schedule int
	Rating             float32
}

type movieDB [MAX_MOV]movie

func (db movieDB) viewAdmin() {
	var input string
	quitApp := 0
	for quitApp == 0 {
		fmt.Println("\n------------------[ CINEZEN ]------------------")
		fmt.Println("\n1. Cari film\n2. Tambah film baru\n3. Edit data film\n4. Tampilkan daftar film\n\n0 untuk keluar")
		fmt.Scan(&input)
		switch input[0] {
		case '0':
			quitApp = 1
		case '1':
			db.cariMovie()
		case '2':
			db.addMovie()
		case '3':
			db.choiceEditMovie()
		case '4':
			db.listMovie()
		default:
			fmt.Println("Input tidak valdi")
		}
	}
}

func (db movieDB) listMovie() {
	fmt.Println("\n----------------[ Daftar Film ]----------------")
	fmt.Printf("\n     %-40s   %-14s   %-4s      %-6s   %-s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal")
	for i := 0; i < db_Len; i++ {
		fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | %-d\n", i+1, db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule)
	}
}

func (db *movieDB) addMovie() {
	var title, genre string
	var duration int
	var schedule int = -1
	var rating float32 = -1

	fmt.Print("\nNama film : ")
	fmt.Scan(&title)
	if len(title) > 40 {
		title = title[0:40]
	}
	fmt.Print("\nGenre film : ")
	fmt.Scan(&genre)
	fmt.Print("\nDurasi film (dalam menit) : ")
	fmt.Scan(&duration)
	fmt.Print("\n")
	for rating == -1 {
		fmt.Print("Rating film : ")
		fmt.Scan(&rating)
		if rating > 10 || rating < 0 {
			rating = -1
			fmt.Println("Rating hanya angka 0-10\n")
		}
	}
	fmt.Print("\n")
	for schedule == -1 {
		fmt.Print("Jadwal film : ")
		fmt.Scan(&schedule)
		if schedule > 20 || schedule < 10 {
			schedule = -1
			fmt.Println("Hanya dapat ditayangkan pukul 10.00 sampai 20.00\n")
		}
	}

	db[db_Len].Title = title
	db[db_Len].Genre = genre
	db[db_Len].Duration = duration
	db[db_Len].Rating = rating
	db[db_Len].Schedule = schedule
	db_Len++
}

func (db *movieDB) cariMovie() {
	var choice string
	var cari interface{}
	found := false
	proceedFind := true

	fmt.Println("\n------------------[ CINEZEN ]------------------")
	fmt.Println("\n1. Cari dengan Judul\n2. Cari dengan Genre\n3. Cari dengan Tanggal\n\n0 Kembali")
	fmt.Scan(&choice)

	switch choice[0] {
	case '1':
		fmt.Print("Masukkan Judul: ")
		var input string
		fmt.Scan(&input)
		cari = input
	case '2':
		fmt.Print("Masukkan Genre: ")
		var input string
		fmt.Scan(&input)
		cari = input
	case '3':
		fmt.Print("Masukkan Jam: ")
		var input int
		fmt.Scan(&input)
		cari = input
	case '0':
		proceedFind = false
	default:
		fmt.Println("\nInput tidak valid")
	}

	if proceedFind {
		i := 0
		for i < db_Len && !found {
			if (choice[0] == '1' && db[i].Title == cari) || (choice[0] == '2' && db[i].Genre == cari) || (choice[0] == '3' && db[i].Schedule == cari) {
				fmt.Println("\n-----------------[ Data Film ]-----------------")
				fmt.Printf("\n%-40s   %-14s   %-4s      %-6s   %s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal")
				fmt.Printf("%-40s | %-14s | %-4d menit | %-6.1f | %-d\n", db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule)
				found = true
				var editChoice string
				fmt.Println("\nApakah Anda ingin mengedit film ini? (y/n)")
				fmt.Scan(&editChoice)
				if editChoice == "y" || editChoice == "Y" {
					db.editMovie(i)
				}
			}
			i++
		}

		if !found {
			fmt.Println("\nFilm tidak ditemukan")
		}
	}
}

func (db *movieDB) choiceEditMovie() {
	var i, dbLen int
	choice := true
	dbLen = (db_Len + 1)
	for choice == true {
		db.listMovie()
		fmt.Println("\nPilih nomor data yang ingin diubah: (0 untuk kembali)")
		fmt.Scan(&i)
		if i > 0 && i < dbLen {
			db.editMovie(i - 1)
		} else if i == 0 {
			fmt.Println("Kembali")
			choice = false
		} else {
			fmt.Println("Pilihan tidak valid.")
			choice = false
		}
	}
}

func (db *movieDB) editMovie(i int) {
	var choice int
	back := true

	for back == true {
		fmt.Println("\nPilih data yang ingin diubah:")
		fmt.Println("1. Judul")
		fmt.Println("2. Genre")
		fmt.Println("3. Durasi")
		fmt.Println("4. Rating")
		fmt.Println("5. Jadwal")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan judul baru: ")
			var newTitle string
			fmt.Scan(&newTitle)
			db[i].Title = newTitle
		case 2:
			fmt.Print("Masukkan genre baru: ")
			var newGenre string
			fmt.Scan(&newGenre)
			db[i].Genre = newGenre
		case 3:
			fmt.Print("Masukkan durasi baru (dalam menit): ")
			var newDuration int
			fmt.Scan(&newDuration)
			db[i].Duration = newDuration
		case 4:
			fmt.Print("Masukkan rating baru: ")
			var newRating float32
			fmt.Scan(&newRating)
			db[i].Rating = newRating
		case 5:
			fmt.Print("Masukkan jadwal baru (misalnya 1400 untuk jam 14:00): ")
			var newSchedule int
			fmt.Scan(&newSchedule)
			db[i].Schedule = newSchedule
		case 0:
			back = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	var db movieDB
	db_Len = 0
	db.viewAdmin()
}
