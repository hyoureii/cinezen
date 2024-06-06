package main

import (
	"cinezen/app"
	"cinezen/db"
	"cinezen/views"
	"fmt"
	"strings"

	"atomicgo.dev/keyboard/keys"
)

func main() {
	var dbMovie db.MovieDB
	var seat db.AvailSeat

	quit := false
	db.Init(&dbMovie)
	for !quit {
		showList := false
		logged := false
		views.ViewStart()
		switch app.DetectKey().String()[0] {
		case '1':
			logged = true
			adminStart(&logged, showList, &dbMovie)
		case '2':
			logged = true
			userStart(&logged, showList, dbMovie, seat)
		case 'q':
			quit = true
		}
	}
}

func adminStart(logged *bool, showList bool, dbMovie *db.MovieDB) {
	for *logged {
		views.ViewAdmin(showList, *dbMovie)
		switch app.DetectKey().String()[0] {
		case '1':
			if showList {
				showList = false
			} else {
				showList = true
			}
		case '2':
			search(*dbMovie)
		case '3':
			edit(dbMovie)
		case '4':
			add(dbMovie)
		case '5':
			delete(dbMovie)
		case '7':
			db.Sort(0, dbMovie)
		case '8':
			db.Sort(1, dbMovie)
		case '9':
			db.Sort(2, dbMovie)
		case 'q':
			*logged = false
		}
	}
}

func userStart(logged *bool, showList bool, dbMovie db.MovieDB, seat db.AvailSeat) {
	showTicket := false
	var dbTicket db.Tickets
	for *logged {
		views.ViewUser(showList, showTicket, dbTicket, dbMovie)
		switch app.DetectKey().String()[0] {
		case '1':
			if showList {
				showList = false
			} else {
				showList = true
			}
		case '2':
			search(dbMovie)
		case '3':
			buyTicket(dbMovie, &dbTicket, &seat)
		case '4':
			if showTicket {
				showTicket = false
			} else {
				showTicket = true
			}
		case '7':
			db.Sort(0, &dbMovie)
		case '8':
			db.Sort(1, &dbMovie)
		case '9':
			db.Sort(2, &dbMovie)
		case 'q':
			*logged = false
		}
	}
}

func search(dbMovie db.MovieDB) db.MovieDB {
	var dbFound db.MovieDB
	var str string
	done := false
	mode := 0
	for !done {
		if len(str) >= 1 {
			switch mode {
			case 0:
				db.SearchTitle(str, dbMovie, &dbFound)
			case 1:
				db.SearchGenre(str, dbMovie, &dbFound)
			case 2:
				temp := strings.Split(str, " ")
				if len(temp) == 3 {
					db.SearchDate(str, dbMovie, &dbFound)
				}
			}
		}
		views.ViewSearch(dbFound, mode, str)
		key := app.DetectKey()
		if key.Code == keys.RuneKey {
			if mode == 1 {
				if key.String()[0] > 47 && key.String()[0] < 58 {
					str = key.String()
				}
			} else if mode == 2 {
				temp := strings.Split(str, " ")
				if len(temp) < 4 {
					str += key.String()
				}
			} else {
				str += key.String()
			}
		} else if key.Code == keys.Backspace {
			if len(str) > 0 {
				str = str[:len(str)-1]
			}
		} else if key.Code == keys.Space {
			if mode == 2 {
				temp := strings.Split(str, " ")
				if len(temp) <= 3 {
					str += " "
				}
			} else {
				str += " "
			}
		} else if key.Code == keys.Tab {
			str = ""
			if mode == 2 {
				mode = 0
			} else {
				mode++
			}
		} else if key.Code == keys.Esc {
			return dbFound
		}
	}
	return dbMovie
}

func add(dbMovie *db.MovieDB) {
	var newMovie db.Movies
	done, success := false, false
	phase := 0
	for !done {
		views.ViewAdd(phase, newMovie, success)
		if success {
			key := app.DetectKey().Code
			if key == keys.Enter {
				newMovie = db.Movies{Title: "", Duration: 0, Genre: "", Rating: 0, Price: 0, Schedule: db.MovieSchedule{Hour: 0, Date: 0, Month: 0}}
				success = false
			} else if key == keys.Esc {
				done = true
			}
		} else {
			switch phase {
			case 0:
				var inp string
				fmt.Scanf("%s\n", &inp)
				inp = strings.ReplaceAll(inp, "_", " ")
				if len(inp) > 0 {
					newMovie.Title = inp
					phase++
				}
			case 1:
				valid := false
				var inp int
				for !valid {
					fmt.Scan(&inp)
					if inp > 0 {
						newMovie.Duration = inp
						valid = true
						phase++
					}
				}
			case 2:
				valid := false
				var inp int
				for !valid {
					fmt.Scan(&inp)
					if inp == 0 {
						inp = 10
					}
					if inp > 0 && inp <= 10 {
						newMovie.Genre = db.GetGenres()[inp-1]
						valid = true
						phase++
					}
				}
			case 3:
				valid := false
				var inp float32
				for !valid {
					fmt.Scan(&inp)
					if inp >= 0 && inp <= 10 {
						newMovie.Rating = inp
						valid = true
						phase++
					}
				}
			case 4:
				valid := false
				var inp int
				for !valid {
					fmt.Scan(&inp)
					if inp >= 30000 {
						newMovie.Price = inp
						valid = true
						phase++
					}
				}
			case 5:
				valid := false
				var inp int
				for !valid {
					fmt.Scan(&inp)
					if inp > 0 && inp < 13 {
						newMovie.Schedule.Month = uint8(inp)
						valid = true
						phase++
					}
				}
			case 6:
				valid := false
				var inp int
				for !valid {
					fmt.Scan(&inp)
					if inp > 0 && inp < 32 {
						newMovie.Schedule.Date = uint8(inp)
						valid = true
						phase++
					}
				}
			case 7:
				valid := false
				var inp int
				for !valid {
					fmt.Scan(&inp)
					if inp > 9 && inp < 22 {
						newMovie.Schedule.Hour = uint8(inp)
						valid = true
						phase++
					}
				}
			case 8:
				key := app.DetectKey().Code
				if key == keys.Tab {
					newMovie = db.Movies{Title: "", Duration: 0, Genre: "", Rating: 0, Price: 0, Schedule: db.MovieSchedule{Hour: 0, Date: 0, Month: 0}}
				} else if key == keys.Enter {
					db.Add(dbMovie, newMovie)
					success = true
				} else if key == keys.Esc {
					done = true
				}
				phase = 0
			}
		}
	}
}

func edit(dbMovie *db.MovieDB) {
	done := false
	idChosen, dataChosen := false, false
	data, id := -1, -1
	for !done {
		doneId := false
		views.ViewEdit(*dbMovie, idChosen, dataChosen, id, data)
		if idChosen {
			for !doneId {
				views.ViewEdit(*dbMovie, idChosen, dataChosen, id, data)
				if dataChosen {
					switch data {
					case 0:
						fmt.Scan(&dbMovie.Db[id].Title)
						dbMovie.Db[id].Title = strings.ReplaceAll(dbMovie.Db[id].Title, "_", " ")
						dataChosen = false
					case 1:
						fmt.Scan(&dbMovie.Db[id].Duration)
						dataChosen = false
					case 2:
						var inp int
						fmt.Scan(&inp)
						if inp == 0 {
							inp = 10
						}
						inp--
						if inp >= 0 && inp < 10 {
							dbMovie.Db[id].Genre = db.GetGenres()[inp]
							dataChosen = false
						}
					case 3:
						var inp float32
						fmt.Scan(&inp)
						if inp >= 0 && inp <= 10 {
							dbMovie.Db[id].Rating = inp
							dataChosen = false
						}
					case 4:
						var inp int
						fmt.Scan(&inp)
						if inp >= 30000 {
							dbMovie.Db[id].Price = inp
							dataChosen = false
						}
					case 5:
						var inp int
						fmt.Scan(&inp)
						if inp > 0 && inp < 13 {
							dbMovie.Db[id].Schedule.Month = uint8(inp)
							data++
						}
					case 6:
						var inp int
						fmt.Scan(&inp)
						if inp > 0 && inp < 32 {
							dbMovie.Db[id].Schedule.Date = uint8(inp)
							data++
						}
					case 7:
						var inp int
						fmt.Scan(&inp)
						if inp > 9 && inp < 22 {
							dbMovie.Db[id].Schedule.Hour = uint8(inp)
							dataChosen = false
						}
					}
				} else {
					var inp string
					fmt.Scan(&inp)
					if inp[0] == 'q' {
						doneId = true
						idChosen = false
					}
					if int(inp[0]-49) > -1 && int(inp[0]-49) < 7 {
						data = int(inp[0] - 49)
						dataChosen = true
					}
				}
			}
		} else {
			var inp int
			fmt.Scan(&inp)
			if inp == 0 {
				done = true
			} else {
				inp--
				if inp > -1 && inp < dbMovie.Len {
					id = inp
					idChosen = true
				}
			}
		}
	}
}

func delete(dbMovie *db.MovieDB) {
	id := -1
	hasChosen, done := false, false
	for !done {
		views.ViewDelete(*dbMovie, id, hasChosen)
		if hasChosen {
			key := app.DetectKey().Code
			if key == keys.Esc {
				id = -1
			} else if key == keys.Enter {
				db.Delete(dbMovie, id)
			}
			hasChosen = false
		} else {
			var inp int
			fmt.Scan(&inp)
			if inp == 0 {
				done = true
			} else {
				inp--
				if inp > -1 && inp < dbMovie.Len {
					id = inp
					hasChosen = true
				}
			}
		}
	}
}

func buyTicket(dbMovie db.MovieDB, dbTicket *db.Tickets, seat *db.AvailSeat) {
	var chosen db.Movies
	id := -1
	done, idChosen, success := false, false, false
	for !done {
		var dbChosen db.MovieDB
		dbChosen.Db[0] = chosen
		dbChosen.Len = 1
		views.BuyTicket(dbMovie, id, idChosen, success)
		if success {
			key := app.DetectKey().Code
			if key == keys.Esc {
				done = true
			} else if key == keys.Enter {
				chosen = db.Movies{Title: "", Duration: 0, Genre: "", Rating: 0, Price: 0, Discount: 0, Schedule: db.MovieSchedule{Hour: 0, Date: 0, Month: 0}}
				success = false
				idChosen = false
			}
		} else {
			if !idChosen {
				views.BuyTicket(dbMovie, id, idChosen, success)
				var inp int
				fmt.Scan(&inp)
				if inp == 0 {
					done = true
				} else {
					inp--
					if inp > -1 && inp < dbMovie.Len {
						id = inp
						idChosen = true
					}
				}
				chosen = dbMovie.Db[id]
				idChosen = true
			} else {
				key := app.DetectKey().Code
				if key == keys.Esc {
					idChosen = false
				} else if key == keys.Enter {
					dbTicket.Movies.Db[dbTicket.Movies.Len] = chosen
					dbTicket.Movies.Len++
					seatChosen := false
					for !seatChosen {
						views.SeatSelect(*seat)
						var seatS string
						var row, column int
						fmt.Scan(&seatS)
						seatS = seatS[:2]
						if (int(seatS[0]-65) > -1 && int(seatS[0]-65) < 10) && (int(seatS[1]-49) > -1 && int(seatS[1]-49) < 10) {
							row = int(seatS[0] - 65)
							column = int(seatS[1] - 49)
						}
						if !seat[row][column] {
							seat[row][column] = true
							dbTicket.Seat = seatS
							seatChosen = true
						}
					}
					success = true
				}
			}
		}
	}
}
