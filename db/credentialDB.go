package db

const MAX_USER = 10

type User struct {
	Username string
	Password string
}

type UserDB [MAX_USER]User
