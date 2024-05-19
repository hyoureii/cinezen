package main

import (
	"fmt"
)

const MAX_MOV int = 100

type movie struct {
	Title, Genre       string
	Duration, Schedule int
	Rating             float32
}

type movieDB struct {
	List [MAX_MOV]movie
	Len  int
}

func viewUser() {
	fmt.Println("\n------------------[ CINEZEN ]------------------")
	fmt.Print("\nCari nama film : ")
	// fmt.Println("-----------------------------------------")
	// printEnter()
}

func (db movieDB) viewAdmin(quit *bool) {
	var input uint8
	fmt.Println("\n--------------[ CINEZEN - ADMIN ]--------------")
	fmt.Println("\n1. Tambah film baru\n2. Edit data film\n3. List film\n\n0 untuk keluar")
	fmt.Scan(&input)
	if input == 0 {
		*quit = true
	} else if input > 3 || input < 0 {
		fmt.Println("Input tidak valid")
	} else if input == 1 {
		db.addMovie()
	} else if input == 3 {
		db.listMovie()
	}
}

func (db movieDB) listMovie() {
	fmt.Println("\n----------------[ Daftar Film ]----------------")
	fmt.Printf("\n     %-20s   %-10s   %-4s   %-4s   %s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal")
	fmt.Println(db.Len)
	for i := 0; i < db.Len; i++ {
		fmt.Printf("%3d. %-20s | %-10s | %-4d | %-4.1f | %-d\n", i+1, db.List[i].Title, db.List[i].Genre, db.List[i].Duration, db.List[i].Rating, db.List[i].Schedule)
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

	//for debugging
	fmt.Println("Judul : ", title)
	fmt.Println("Genre : ", genre)
	fmt.Println("Durasi : ", duration)
	fmt.Println("Rating : ", rating)
	fmt.Println("Jadwal : ", schedule)

	db.List[db.Len].Title = title
	db.List[db.Len].Genre = genre
	db.List[db.Len].Duration = duration
	db.List[db.Len].Rating = rating
	db.List[db.Len].Schedule = schedule
	db.Len++
}

func main() {
	var db movieDB
	db.Len = 0
	quitApp := false
	for !quitApp {
		db.viewAdmin(&quitApp)
	}
}
