package main

import (
	"comic-rss-backend/db"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var seedUsers = []struct {
	UserName string
	Password string
}{
	{UserName: "saki", Password: "password123"},
	{UserName: "taro", Password: "password123"},
}

func main() {
	database, err := db.Open("comic_rss.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	for _, u := range seedUsers {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash password for %s: %v", u.UserName, err)
		}

		_, err = database.Exec(
			`INSERT INTO users (user_name, password_hash) VALUES (?, ?)
			 ON CONFLICT(user_name) DO UPDATE SET password_hash = excluded.password_hash`,
			u.UserName, string(hash),
		)
		if err != nil {
			log.Fatalf("failed to seed user %s: %v", u.UserName, err)
		}

		log.Printf("seeded user: %s\n", u.UserName)
	}
}
