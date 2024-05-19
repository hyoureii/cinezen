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
	var input byte
	quitApp := 0
	for quitApp == 0 {
		fmt.Println("\n------------------[ CINEZEN ]------------------")
		fmt.Println("\n1. Cari film\n2. Tambah film baru\n3. Edit data film\n4. Tampilkan daftar film\n\n0 untuk keluar")
		fmt.Scanf("%c\n", &input)
		switch input {
		case '0':
			quitApp = 1
		case '1':
			db.cariMovie()
		case '2':
			db.addMovie()
		case '4':
			db.listMovie()
		default:
			fmt.Println("Input tidak valdi")
		}
	}
}

func (db movieDB) listMovie() {
	fmt.Println("\n----------------[ Daftar Film ]----------------")
	fmt.Printf("\n     %-20s   %-10s   %-4s   %-4s   %s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal")
	for i := 0; i < db_Len; i++ {
		fmt.Printf("%3d. %-20s | %-10s | %-4d | %-4.1f | %-d\n", i+1, db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule)
	}
}

func (db *movieDB) addMovie() {
	var title, genre string
	var duration int
	var schedule int = -1
	var rating float32 = -1

	fmt.Print("\nNama film : ")
	fmt.Scan(&title)
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

func (db movieDB) cariMovie() {
	var choice byte
	var cari interface{}
	found := false

	fmt.Println("\n------------------[ CINEZEN ]------------------")
	fmt.Println("\n1. Cari dengan Judul\n2. Cari dengan Genre\n3. Cari dengan Tanggal\n\n0 Kembali")
	fmt.Scanf("%c\n", &choice)

	switch choice {
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
		db.viewAdmin()
	default:
		fmt.Println("Input tidak valid")
		db.cariMovie()
	}

	for _, movie := range db {
		if (choice == '1' && movie.Title == cari) ||
			(choice == '2' && movie.Genre == cari) ||
			(choice == '3' && movie.Schedule == cari) {
			fmt.Println("\n----------------[ Daftar Film ]----------------")
			fmt.Printf("\n%-20s   %-10s   %-4s   %-4s   %s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal")
			fmt.Printf("%-20s | %-10s | %-4d | %-4.1f | %-d\n", movie.Title, movie.Genre, movie.Duration, movie.Rating, movie.Schedule)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Film tidak ditemukan")
	}
}

func main() {
	var db movieDB
	db_Len = 0
	db.viewAdmin()
}
