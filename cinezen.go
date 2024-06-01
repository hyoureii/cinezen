package main

import (
	"cinezen/app"
	"cinezen/db"
	"cinezen/views"
)

func main() {
	var dbMovie db.MovieDB
	// var dbUser db.UserDB

	quit := false
	db.Init(&dbMovie)
	for !quit {
		showList := false
		logged := false
		views.ViewStart()
		switch app.DetectKey().String()[0] {
		case '1':
			logged = true
			for logged {
				views.ViewAdmin(showList, dbMovie)
				switch app.DetectKey().String()[0] {
				case '1':
					if showList {
						showList = false
					} else {
						showList = true
					}
				case '2':
					views.ViewSearch(dbMovie)
					switch app.DetectKey().String()[0] {
					case 'q':
					}
				}
			}
		case '2':
			logged = true
			for logged {
				views.ViewUser(showList, dbMovie)
				switch app.DetectKey().String()[0] {
				case '1':
					if showList {
						showList = false
					} else {
						showList = true
					}
				}
			}
		}
	}
}
