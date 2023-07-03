package main

import (
	"github.com/LuisKpBeta/url-shortener/internal/api"
	"github.com/LuisKpBeta/url-shortener/internal/database"
	"github.com/LuisKpBeta/url-shortener/internal/token"
	"github.com/LuisKpBeta/url-shortener/pk/repository"
	service_url "github.com/LuisKpBeta/url-shortener/pk/service/url"
)

func main() {
	conn := database.ConnectToDatabase()
	createAction := repository.CreateUrl(conn)
	parameters := service_url.CreateParameters{
		CreateToken:       token.GenerateUrlToken,
		SaveUrlRepository: createAction,
	}

	server := api.CreateHttpServer()
	api.CreateUrlShortnerHandler(server, service_url.Create(parameters))
	api.StartHttpServer(server)

}
