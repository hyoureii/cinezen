package db

const MAX_MOV int = 100

type MovieSchedule struct {
	Hour  uint8
	Day   uint8
	Month uint8
}

type Movies struct {
	Title    string
	Duration int
	Genre    string
	Rating   float32
	Price    int
	Schedule MovieSchedule
}

type MovieDB struct {
	Db  [MAX_MOV]Movies
	Len int
}

// InitDB buat ngefill database film, buat ngetest fitur lain
func Init(db *MovieDB) {
	db.Db[0] = Movies{"The Shawshank Redemption", 142, "Drama", 9.3, 40000, MovieSchedule{10, 1, 1}}
	db.Db[1] = Movies{"The Godfather", 175, "Drama", 9.2, 48000, MovieSchedule{14, 2, 2}}
	db.Db[2] = Movies{"Pulp Fiction", 154, "Thriller", 8.9, 60000, MovieSchedule{18, 3, 3}}
	db.Db[3] = Movies{"The Dark Knight", 152, "Action", 9.0, 72000, MovieSchedule{20, 4, 4}}
	db.Db[4] = Movies{"Inception", 148, "Sci-Fi", 8.8, 55000, MovieSchedule{22, 5, 5}}
	db.Len = 5
}

func Add(db *MovieDB) {

}
