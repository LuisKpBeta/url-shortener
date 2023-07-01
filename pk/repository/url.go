package repository

import (
	"database/sql"

	"github.com/LuisKpBeta/url-shortener/pk/entity"
)

func CreateUrl(db *sql.DB) func(*entity.Url) error {
	return func(url *entity.Url) error {
		stmt, err := db.Prepare("INSERT INTO urls (original, shortened, created_at) values(?,?,?)")
		checkErr(err)
		result, err := stmt.Exec(url.Original, url.Shortened, url.CreatedAt)
		if err != nil {
			return err
		}
		createdId, err := result.LastInsertId()
		checkErr(err)
		url.Id = int(createdId)
		return nil
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
