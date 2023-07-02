package repository

import (
	"database/sql"

	"github.com/LuisKpBeta/url-shortener/pk/entity"
)

func CreateUrl(db *sql.DB) func(*entity.Url) error {
	return func(url *entity.Url) error {
		// stmt, err := db.Prepare("INSERT INTO url (original, shortened, createdat) VALUES (? , ? , ?)")
		stmt, err := db.Prepare("INSERT INTO url (original, shortened, createdat) VALUES ($1, $2, $3) RETURNING id")
		checkErr(err)
		defer stmt.Close()
		var createdId int
		stmt.QueryRow(url.Original, url.Shortened, url.CreatedAt).Scan(createdId)
		url.Id = int(createdId)
		return nil
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
