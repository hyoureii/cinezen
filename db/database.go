package db

import (
	"cinezen/app"
	"fmt"
	"strings"
)

const MAX_MOV int = 1000

type MovieSchedule struct {
	Hour  uint8
	Date  uint8
	Month uint8
}

type Movies struct {
	Title    string
	Duration int
	Genre    string
	Rating   float32
	Price    int
	Discount int
	Schedule MovieSchedule
}

type MovieDB struct {
	Db  [MAX_MOV]Movies
	Len int
}

type AvailSeat [9][9]bool

type Tickets struct {
	Movies MovieDB
	Seat   string
}

func ConvertMonth(month uint8) string {
	switch month {
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

func GetGenres() [10]string {
	return [10]string{"Action", "Comedy", "Drama", "Horror", "Romance", "Sci-Fi", "Documentary", "Thriller", "Animation", "Fantasy"}
}

func CheckGenre(str string) int {
	switch str {
	case "Action":
		return 0
	case "Comedy":
		return 1
	case "Drama":
		return 2
	case "Horror":
		return 3
	case "Romance":
		return 4
	case "Sci-Fi":
		return 5
	case "Documentary":
		return 6
	case "Thriller":
		return 7
	case "Animation":
		return 8
	case "Fantasy":
		return 9
	}
	return -1
}

// InitDB buat ngefill database film, buat ngetest fitur lain
func Init(db *MovieDB) {
	db.Db[0] = Movies{"The Shawshank Redemption", 142, "Drama", 9.3, 40000, app.GenerateDiscount(), MovieSchedule{10, 1, 1}}
	db.Db[1] = Movies{"The Godfather", 175, "Drama", 9.2, 48000, app.GenerateDiscount(), MovieSchedule{14, 2, 2}}
	db.Db[2] = Movies{"Pulp Fiction", 154, "Thriller", 8.9, 60000, app.GenerateDiscount(), MovieSchedule{18, 3, 3}}
	db.Db[3] = Movies{"The Dark Knight", 152, "Action", 9.0, 72000, app.GenerateDiscount(), MovieSchedule{20, 4, 4}}
	db.Db[4] = Movies{"Inception", 148, "Sci-Fi", 8.8, 55000, app.GenerateDiscount(), MovieSchedule{13, 5, 5}}
	db.Db[5] = Movies{"The Lord of the Rings: The Return of the King", 251, "Fantasy", 8.9, 90000, app.GenerateDiscount(), MovieSchedule{10, 6, 6}}
	db.Db[6] = Movies{"The Lord of the Rings: The Fellowship of the Ring", 178, "Fantasy", 8.8, 70000, app.GenerateDiscount(), MovieSchedule{14, 7, 7}}
	db.Db[7] = Movies{"The Lord of the Rings: The Two Towers", 179, "Fantasy", 8.7, 70000, app.GenerateDiscount(), MovieSchedule{18, 8, 8}}
	db.Db[8] = Movies{"The Matrix", 136, "Sci-Fi", 8.7, 50000, app.GenerateDiscount(), MovieSchedule{20, 9, 9}}
	db.Db[9] = Movies{"The Matrix Reloaded", 138, "Sci-Fi", 8.6, 50000, app.GenerateDiscount(), MovieSchedule{13, 10, 10}}
	db.Len = 10
}

func Add(db *MovieDB, stats Movies) {
	db.Db[db.Len] = stats
	db.Len++
}

func SearchTitle(str string, from MovieDB, to *MovieDB) {
	to.Len = 0
	for i := 0; i < from.Len; i++ {
		if strings.Contains(strings.ToLower(from.Db[i].Title), strings.ToLower(str)) {
			to.Db[to.Len] = from.Db[i]
			to.Len++
		}
	}
}

func SearchGenre(str string, from MovieDB, to *MovieDB) {
	g := int(str[0]-48) - 1
	if g == -1 {
		g = 9
	}
	to.Len = 0
	for i := 0; i < from.Len; i++ {
		if strings.EqualFold(from.Db[i].Genre, GetGenres()[g]) {
			to.Db[to.Len] = from.Db[i]
			to.Len++
		}
	}
}

func SearchDate(str string, from MovieDB, to *MovieDB) {
	date := strings.Split(str, " ")
	to.Len = 0
	for i := 0; i < from.Len; i++ {
		if date[0] == fmt.Sprintf("%d", from.Db[i].Schedule.Hour) && date[1] == fmt.Sprintf("%d", from.Db[i].Schedule.Date) && date[2] == fmt.Sprintf("%d", from.Db[i].Schedule.Month) {
			to.Db[to.Len] = from.Db[i]
			to.Len++
		}
	}
}

func Sort(mode int, db *MovieDB) {
	switch mode {
	case 0:
		for i := 1; i < db.Len; i++ {
			j := i
			for j > 0 {
				if CheckGenre(db.Db[j-1].Genre) > CheckGenre(db.Db[j].Genre) {
					db.Db[j], db.Db[j-1] = db.Db[j-1], db.Db[j]
				}
				j--
			}
		}
	case 1:
		for i := 1; i < db.Len; i++ {
			j := i
			for j > 0 {
				if db.Db[j-1].Schedule.Month > db.Db[j].Schedule.Month {
					db.Db[j], db.Db[j-1] = db.Db[j-1], db.Db[j]
				} else if db.Db[j-1].Schedule.Month == db.Db[j].Schedule.Month {
					if db.Db[j-1].Schedule.Date > db.Db[j].Schedule.Date {
						db.Db[j], db.Db[j-1] = db.Db[j-1], db.Db[j]
					} else if db.Db[j-1].Schedule.Date == db.Db[j].Schedule.Date {
						if db.Db[j-1].Schedule.Hour > db.Db[j].Schedule.Hour {
							db.Db[j], db.Db[j-1] = db.Db[j-1], db.Db[j]
						}
					}
				}
				j--
			}
		}
	case 2:
		for i := 1; i < db.Len; i++ {
			j := i
			for j > 0 {
				if db.Db[j-1].Price > db.Db[j].Price {
					db.Db[j], db.Db[j-1] = db.Db[j-1], db.Db[j]
				}
				j--
			}
		}
	}
}

func Delete(db *MovieDB, id int) {
	for i := id; i < db.Len-1; i++ {
		db.Db[i] = db.Db[i+1]
	}
	db.Len--
}
