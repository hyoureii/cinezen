package app

var (
	IsAdmin    bool
	AppState   string
	WhatToShow string
)

func InitApp() {
	IsAdmin = false
	AppState = "notLogged"
	WhatToShow = "choices"
}

func UpdateApp(cursor int) {
	switch AppState {
	case "notLogged":
		if cursor == 1 {
			AppState = "mainAdmin"
		}
	case "mainAdmin":
		switch cursor {
		case 3:
			AppState = "notLogged"
		}
	}
}
