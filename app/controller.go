package app

var (
	IsAdmin  bool
	AppState string
)

func InitApp() {
	AppState = "notLogged"
}

func UpdateApp(cursor int) {
	switch AppState {
	case "notLogged":

	}
}
