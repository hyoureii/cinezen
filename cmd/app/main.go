package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	MAX_MOV      int = 100
	maxPurchases int = 100
)

var db_Len int

type movie struct {
	Title, Genre    string
	Duration, Price int
	Rating          float32
	Discount        Discount
	Global          Global
	Schedule        MovieSchedule
}

type Discount struct {
	Percentage float64
	MinPrice   int
}

type Purchase struct {
	Title, Genre    string
	Duration, Price int
	Rating          float32
	Discount        Discount
	Schedule        MovieSchedule
}

type Global struct {
	Authority     bool
	purchaseCount int
	purchases     [maxPurchases]Purchase
}

type MovieSchedule struct {
	Hour  int
	Day   int
	Month int
}

type movieDB [MAX_MOV]movie

var clearScreen func() = func() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func genRanDiscount(price int, discount *Discount) {
	var seed int64
	for i := 0; i < 10; i++ {
		seed += time.Now().UnixNano()
	}
	rand.Seed(seed)
	if price >= 40000 && price <= 100000 {
		discount.Percentage = rand.Float64() * 0.05
	} else if price > 32000 && price < 40000 {
		discount.Percentage = rand.Float64() * 0.01
	} else if price > 100000 {
		discount.Percentage = rand.Float64() * 0.10
	} else {
		discount.Percentage = 0
	}
	discount.MinPrice = price - int(float64(price)*discount.Percentage)
}

func (db *movieDB) viewAdmin() {
	var input string
	quitApp := 0
	for quitApp == 0 {
		fmt.Println("\n------------------[ CINEZEN ]------------------")
		fmt.Println("\n1. Cari film\n2. Tambah film baru\n3. Edit data film\n4. Tampilkan daftar film\n\n0 untuk keluar")
		fmt.Scan(&input)
		switch input[0] {
		case '0':
			clearScreen()
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
			clearScreen()
			fmt.Println("Input tidak valid.")
		}
	}
}

func (db *movieDB) viewUser() {
	var input string
	quitApp := 0
	for quitApp == 0 {
		fmt.Println("\n------------------[ CINEZEN ]------------------")
		fmt.Println("\n1. List Film\n2. Cari Film\n3. List Pembelian\n\n0 untuk keluar")
		fmt.Scan(&input)
		switch input[0] {
		case '0':
			clearScreen()
			quitApp = 1
		case '1':
			clearScreen()
			db.listMovie()
		case '2':
			clearScreen()
			db.cariMovie()
		case '3':
			clearScreen()
			db.purcList()
		default:
			clearScreen()
			fmt.Println("Input tidak valid.")
		}
	}
}

func (db movieDB) listMovie() {
	fmt.Println("\n----------------[ Daftar Film ]----------------")
	fmt.Printf("\n     %-40s   %-14s   %-4s      %-6s          %-12s          %4s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal", "Harga")
	for i := 0; i < db_Len; i++ {
		if db[i].Discount.Percentage*100 > 0 {
			fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | Pukul: %-2d Tanggal: %-2d %-2d | Rp.%4d	| Diskon %.1f%% (Total Harga: Rp.%d)\n", i+1, db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule.Hour, db[i].Schedule.Day, db[i].Schedule.Month, db[i].Price, db[i].Discount.Percentage*100, db[i].Discount.MinPrice)
		} else {
			fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | Pukul: %-2d Tanggal: %-2d %-2d | Rp.%4d	| Tidak Dapat Diskon\n", i+1, db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule.Hour, db[i].Schedule.Day, db[i].Schedule.Month, db[i].Price)
		}
	}
	if !db[0].Global.Authority {
		db.beliTiket()
	}
}

func (db *movieDB) addMovie() {
	var (
		title, chosenGenre string
		duration           int     = -1
		scheduleHour       int     = -1
		scheduleDay        int     = -1
		scheduleMonth      int     = -1
		rating             float32 = -1
		price              int     = -1
		discount           Discount
		choice             bool
	)
	choice = true

	fmt.Print("\nNama film : ")
	fmt.Scan(&title)
	if len(title) > 40 {
		title = title[0:40]
	}

	for choice {
		var genre string
		fmt.Printf("\nPilih Genre: (1-9)")
		fmt.Printf("\n1. Action\n2. Comedy\n3. Drama\n4. Horor\n5. Romance\n6. Sci-Fi\n7. Documentary\n8. Thriller\n9. Animation\n\n0. Kembali\n")
		fmt.Scan(&genre)
		switch genre[0] {
		case '1':
			chosenGenre = "Action"
			choice = false
		case '2':
			chosenGenre = "Comedy"
			choice = false
		case '3':
			chosenGenre = "Drama"
			choice = false
		case '4':
			chosenGenre = "Horror"
			choice = false
		case '5':
			chosenGenre = "Romance"
			choice = false
		case '6':
			chosenGenre = "Sci-Fi"
			choice = false
		case '7':
			chosenGenre = "Documentary"
			choice = false
		case '8':
			chosenGenre = "Thriller"
			choice = false
		case '9':
			chosenGenre = "Animation"
			choice = false
		case '0':
			choice = false
		}
	}

	fmt.Print("\n")
	for duration == -1 {
		fmt.Print("Duration film : ")
		fmt.Scan(&duration)
		if duration < 1 {
			duration = -1
			fmt.Printf("\nDurasi minimal 1 menit\n")
		}
	}

	fmt.Print("\n")
	for rating == -1 {
		fmt.Print("Rating film : ")
		fmt.Scan(&rating)
		if rating > 10 || rating < 0 {
			rating = -1
			fmt.Printf("\nRating hanya angka 0-10\n")
		}
	}
	fmt.Print("\n")
	for scheduleHour == -1 || scheduleDay == -1 || scheduleMonth == -1 {
		fmt.Print("Jadwal film tayang : (format: Hour Day Month)")
		fmt.Scanf("%d %d %d", scheduleHour, scheduleDay, scheduleMonth)
		if scheduleHour > 20 || scheduleHour < 10 {
			scheduleHour = -1
			fmt.Printf("\nHanya dapat ditayangkan pukul 10 sampai 20\n")
		}
		if scheduleDay > 0 || scheduleDay < 32 {
			scheduleDay = -1
			fmt.Printf("\nHanya dapat ditayangkan tanggal 1 sampai 31\n")
		}
		if scheduleMonth > 0 || scheduleMonth < 13 {
			scheduleMonth = -1
			fmt.Printf("\nHanya dapat ditayangkan bulan 1 sampai 12\n")
		}
	}
	fmt.Print("\n")
	for price == -1 {
		fmt.Print("Harga Tiket : ")
		fmt.Scan(&price)
		if price < 30000 {
			price = -1
			fmt.Printf("\nMinimal Harga Rp.30000\n")
		}
	}

	genRanDiscount(price, &discount)

	db[db_Len].Title = title
	db[db_Len].Genre = chosenGenre
	db[db_Len].Duration = duration
	db[db_Len].Rating = rating
	db[db_Len].Schedule.Hour = scheduleHour
	db[db_Len].Schedule.Day = scheduleDay
	db[db_Len].Schedule.Month = scheduleMonth
	db[db_Len].Price = price
	db[db_Len].Discount = discount
	db_Len++
}

func (db *movieDB) cariMovie() {
	var (
		choice string
		cari   interface{}
		choose bool
	)
	choose = true
	found := false
	proceedFind := true

	fmt.Println("\n------------------[ CINEZEN ]------------------")
	fmt.Println("\n1. Cari dengan Judul\n2. Cari dengan Genre\n3. Cari dengan Jadwal\n\n0 Kembali")
	fmt.Scan(&choice)

	switch choice[0] {
	case '1':
		fmt.Print("Masukkan Judul: ")
		var input string
		fmt.Scan(&input)
		cari = input
	case '2':
		for choose {
			var genre string
			fmt.Printf("\nPilih Genre: (1-9)")
			fmt.Printf("\n1. Action\n2. Comedy\n3. Drama\n4. Horor\n5. Romance\n6. Sci-Fi\n7. Documentary\n8. Thriller\n9. Animation\n\n0. Kembali\n")
			fmt.Scan(&genre)
			switch genre[0] {
			case '1':
				cari = "Action"
				choose = false
			case '2':
				cari = "Comedy"
				choose = false
			case '3':
				cari = "Drama"
				choose = false
			case '4':
				cari = "Horror"
				choose = false
			case '5':
				cari = "Romance"
				choose = false
			case '6':
				cari = "Sci-Fi"
				choose = false
			case '7':
				cari = "Documentary"
				choose = false
			case '8':
				cari = "Thriller"
				choose = false
			case '9':
				cari = "Animation"
				choose = false
			default:
				fmt.Print("\nPilihan tidak valid.\n")
			}
		}
	case '3':
		fmt.Print("Masukkan Jadwal: (10-20)")
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
		fmt.Printf("\n----------------[ Data Film %s ]----------------", cari)
		fmt.Printf("\n     %-40s   %-14s   %-4s      %-6s   %-4s    %-s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal", "Harga")
		for i < db_Len && !found {
			if (choice[0] == '1' && db[i].Title == cari) || (choice[0] == '2' && db[i].Genre == cari) || (choice[0] == '3' && db[i].Schedule == cari) {
				if db[i].Discount.Percentage*100 > 0.0000000000000 {
					fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | %-6.d | Rp.%-d	| Diskon %.1f%% (Total Harga: Rp.%d)\n", i+1, db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule, db[i].Price, db[i].Discount.Percentage*100, db[i].Discount.MinPrice)

				} else {
					fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | %-6.d | Rp.%-d	| Tidak Dapat Diskon\n", i+1, db[i].Title, db[i].Genre, db[i].Duration, db[i].Rating, db[i].Schedule, db[i].Price)
				}
				if (i + 1) == db_Len {
					found = true
				}
			}
			i++
		}
		if db[0].Global.Authority {
			fmt.Println("Apakah Anda Ingin Mengedit Data?")
			var choose string
			fmt.Scan(&choose)
			if choose == "y" || choose == "Y" {
				db.editMovie(i)
			}
		}
		if !db[0].Global.Authority {
			db.beliTiket()
		}

		if !found {
			fmt.Println("\nFilm tidak ditemukan")
		}
	}
}

func (db *movieDB) choiceEditMovie() {
	var dbLen int
	var j string
	choice := true
	dbLen = (db_Len + 1)
	for choice {
		db.listMovie()
		fmt.Println("\nPilih nomor data yang ingin diubah: (0 untuk kembali)")
		fmt.Scan(&j)
		i := int(j[0]) - 48
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

	for back {
		fmt.Println("\nPilih data yang ingin diubah:")
		fmt.Printf("1. Judul\n2. Genre\n3. Durasi\n4. Rating\n5. Jadwal\n6. Harga\n\n0. Kembali")
		fmt.Print("\nPilihan: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan judul baru: ")
			var newTitle string
			fmt.Scan(&newTitle)
			if len(newTitle) > 40 {
				newTitle = newTitle[0:40]
			}
			db[i].Title = newTitle
		case 2:
			choose := true
			for choose {
				var newGenre, genre string
				fmt.Printf("\nPilih Genre: (1-9)")
				fmt.Printf("\n1. Action\n2. Comedy\n3. Drama\n4. Horor\n5. Romance\n6. Sci-Fi\n7. Documentary\n8. Thriller\n9. Animation\n\n")
				fmt.Scan(&genre)
				switch genre[0] {
				case '1':
					newGenre = "Action"
					db[i].Genre = newGenre
					choose = false
				case '2':
					newGenre = "Comedy"
					db[i].Genre = newGenre
					choose = false
				case '3':
					newGenre = "Drama"
					db[i].Genre = newGenre
					choose = false
				case '4':
					newGenre = "Horror"
					db[i].Genre = newGenre
					choose = false
				case '5':
					newGenre = "Romance"
					db[i].Genre = newGenre
					choose = false
				case '6':
					newGenre = "Sci-Fi"
					db[i].Genre = newGenre
					choose = false
				case '7':
					newGenre = "Documentary"
					db[i].Genre = newGenre
					choose = false
				case '8':
					newGenre = "Thriller"
					db[i].Genre = newGenre
					choose = false
				case '9':
					newGenre = "Animation"
					db[i].Genre = newGenre
					choose = false
				default:
					fmt.Print("\nPilihan tidak valid.\n")
				}
			}
		case 3:
			var newDuration int = -1
			for newDuration == -1 {
				fmt.Print("Masukkan durasi film baru : ")
				fmt.Scan(&newDuration)
				if newDuration < 1 {
					newDuration = -1
					fmt.Printf("\nDurasi minimal 1 menit\n")
				}
			}
		case 4:
			var newRating float32 = -1
			for newRating == -1 {
				fmt.Print("Masukkan rating baru: ")
				fmt.Scan(&newRating)
				if newRating > 10 || newRating < 0 {
					newRating = -1
					fmt.Printf("\nRating hanya angka 0-10\n")
				}
			}
			db[i].Rating = newRating
		case 5:
			var (
				newScheduleHour  int = -1
				newScheduleDay   int = -1
				newScheduleMonth int = -1
			)
			for newScheduleHour == -1 || newScheduleDay == -1 || newScheduleMonth == -1 {
				fmt.Print("Jadwal film tayang : (format: Hour Day Month)")
				fmt.Scanf("%d %d %d", newScheduleHour, newScheduleDay, newScheduleMonth)
				if newScheduleHour > 20 || newScheduleHour < 10 {
					newScheduleHour = -1
					fmt.Printf("\nHanya dapat ditayangkan pukul 10 sampai 20\n")
				}
				if newScheduleDay > 0 || newScheduleDay < 32 {
					newScheduleDay = -1
					fmt.Printf("\nHanya dapat ditayangkan tanggal 1 sampai 31\n")
				}
				if newScheduleMonth > 0 || newScheduleMonth < 13 {
					newScheduleMonth = -1
					fmt.Printf("\nHanya dapat ditayangkan bulan 1 sampai 12\n")
				}
			}
			db[i].Schedule.Hour = newScheduleHour
			db[i].Schedule.Day = newScheduleDay
			db[i].Schedule.Month = newScheduleMonth

		case 6:
			var newPrice int = -1
			for newPrice == -1 {
				fmt.Print("Masukkan Harga tiket baru: ")
				fmt.Scan(&newPrice)
				if newPrice < 15000 {
					newPrice = -1
					fmt.Printf("\nMinimal Harga tiket Rp.15000\n")
				}
			}
			db[i].Price = newPrice
		case 0:
			back = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func (db *movieDB) purcList() {
	fmt.Println("\n----------------[ Tiket Yang Dimiliki ]----------------")
	fmt.Printf("\n     %-40s   %-14s   %-4s      %-6s          %-12s          %4s\n", "Judul", "Genre", "Durasi", "Rating", "Jadwal", "Harga")
	for i := 0; i < db[0].Global.purchaseCount; i++ {
		p := db[0].Global.purchases[i]
		if p.Discount.Percentage*100 > 0.0000000000000 {
			fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | Pukul: %-2d Tanggal: %-2d %-2d | Rp.%-d\t	|\n", i+1, p.Title, p.Genre, p.Duration, p.Rating, p.Schedule.Hour, p.Schedule.Day, p.Schedule.Month, p.Discount.MinPrice)
		} else {
			fmt.Printf("%3d. %-40s | %-14s | %-4d menit | %-6.1f | Pukul: %-2d Tanggal: %-2d %-2d | Rp.%-d\t	|\n", i+1, p.Title, p.Genre, p.Duration, p.Rating, p.Schedule.Hour, p.Schedule.Day, p.Schedule.Month, p.Discount.MinPrice)
		}
	}
}

func (db *movieDB) beliTiket() {
	var (
		choice string
		cari   string
		found  bool
	)

	fmt.Println("\n------------------[ CINEZEN ]------------------")
	fmt.Println("\n1. Beli Tiket\n\n0 Kembali")
	fmt.Scan(&choice)

	switch choice[0] {
	case '1':
		fmt.Print("Masukkan Judul Film: ")
		fmt.Scan(&cari)
	case '0':
		return
	default:
		fmt.Println("\nInput tidak valid")
		return
	}

	found = false
	for i := 0; i < db_Len; i++ {
		if db[i].Title == cari {
			if db[0].Global.purchaseCount < maxPurchases {
				db[0].Global.purchases[db[0].Global.purchaseCount] = Purchase{
					Title:    db[i].Title,
					Genre:    db[i].Genre,
					Duration: db[i].Duration,
					Schedule: MovieSchedule{db[i].Schedule.Hour, db[i].Schedule.Day, db[i].Schedule.Month},
					Price:    db[i].Price,
					Rating:   db[i].Rating,
					Discount: db[i].Discount,
				}
				db[0].Global.purchaseCount++
				found = true

				fmt.Printf("\nTiket untuk film '%s' berhasil dibeli!\n", db[i].Title)
			} else {
				fmt.Println("\nDaftar pembelian sudah penuh")
			}
			break
		}
	}

	if !found {
		fmt.Println("\nFilm tidak ditemukan")
	}
}

func main() {
	var db movieDB
	var choice string
	db[0].Global.Authority = false
	db[0].Global.purchaseCount = 0
	db_Len = 0

	db[db_Len] = movie{
		Title:    "Inception",
		Genre:    "Fantasy",
		Duration: 148,
		Schedule: MovieSchedule{14, 12, 10},
		Rating:   8.8,
		Price:    30000,
	}
	db_Len++
	db[db_Len] = movie{
		Title:    "BadGuy",
		Genre:    "Fantasy",
		Duration: 200,
		Schedule: MovieSchedule{10, 28, 6},
		Rating:   9.8,
		Price:    45000,
	}
	db_Len++

	db[db_Len] = movie{
		Title:    "Spy",
		Genre:    "Action",
		Duration: 128,
		Schedule: MovieSchedule{10, 7, 2},
		Rating:   7.4,
		Price:    35000,
	}
	db_Len++
	db[db_Len] = movie{
		Title:    "Stressed",
		Genre:    "Romance",
		Duration: 174,
		Schedule: MovieSchedule{17, 2, 6},
		Rating:   9.2,
		Price:    50000,
	}
	db_Len++
	db[db_Len] = movie{
		Title:    "Hunter",
		Genre:    "Action",
		Duration: 243,
		Schedule: MovieSchedule{19, 28, 5},
		Rating:   8.5,
		Price:    200000,
	}
	db_Len++

	for i := 0; i < db_Len; i++ {
		genRanDiscount(db[i].Price, &db[i].Discount)
	}

	quitApp := 0
	for quitApp == 0 {
		clearScreen()
		fmt.Println("\n------------------[ CINEZEN ]------------------")
		fmt.Println("Login Dulu Lah Zayang")
		fmt.Println("\n1. Admin\n2. User\n\n0 untuk keluar")
		fmt.Scan(&choice)
		switch choice[0] {
		case '0':
			quitApp = 1
		case '1':
			db[0].Global.Authority = true
			db.viewAdmin()
		case '2':
			db[0].Global.Authority = false
			db.viewUser()
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
