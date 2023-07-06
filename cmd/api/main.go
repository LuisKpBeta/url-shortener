package main

import (
	"log"

	"github.com/LuisKpBeta/url-shortener/internal/api"
	"github.com/LuisKpBeta/url-shortener/internal/database"
	"github.com/LuisKpBeta/url-shortener/internal/prometheus"
	"github.com/LuisKpBeta/url-shortener/internal/token"
	"github.com/LuisKpBeta/url-shortener/pk/repository"
	service_url "github.com/LuisKpBeta/url-shortener/pk/service/url"
)

func main() {
	conn := database.ConnectToDatabase()

	prometheusService, err := prometheus.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	createAction := repository.CreateUrl(conn)
	parameters := service_url.CreateParameters{
		CreateToken:       token.GenerateUrlToken,
		SaveUrlRepository: createAction,
	}

	readToken := repository.GetUrlByToken(conn)
	readTokenParameters := service_url.ReadUrlParameters{
		GetUrlRepository: readToken,
	}

	server := api.CreateHttpServer()
	api.CreatePrometheusHandler(server)
	api.CreateMetricsMiddleware(server, prometheusService)
	api.GetUrlByTokenHandler(server, service_url.ReadUrl(readTokenParameters))
	api.CreateUrlShortnerHandler(server, service_url.Create(parameters))

	api.StartHttpServer(server)

}
