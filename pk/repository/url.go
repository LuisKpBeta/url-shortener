package repository

import (
	"database/sql"

	"github.com/LuisKpBeta/url-shortener/pk/entity"
	service_url "github.com/LuisKpBeta/url-shortener/pk/service/url"
)

func CreateUrl(db *sql.DB) func(*entity.Url) error {
	return func(url *entity.Url) error {
		stmt, err := db.Prepare("INSERT INTO url (original, shortened, createdat) VALUES ($1, $2, $3) RETURNING id")
		checkErr(err)
		defer stmt.Close()
		var createdId int
		stmt.QueryRow(url.Original, url.Shortened, url.CreatedAt).Scan(&createdId)
		url.Id = int(createdId)
		return nil
	}
}

func GetUrlByToken(db *sql.DB) func(string) (service_url.ReadUrlData, error) {
	return func(token string) (service_url.ReadUrlData, error) {
		var url string
		var id int
		err := db.QueryRow("SELECT original, id FROM url where shortened = $1", token).
			Scan(&url, &id)
		var urlData service_url.ReadUrlData
		if err != nil {
			if err == sql.ErrNoRows {
				return urlData, nil
			}
			return urlData, err
		}
		urlData.Id = id
		urlData.Original = url
		return urlData, nil
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
