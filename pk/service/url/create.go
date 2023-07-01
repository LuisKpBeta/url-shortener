package service_url

import (
	"time"

	"github.com/LuisKpBeta/url-shortener/pk/entity"
)

type CreateParameters struct {
	CreateToken       func() string
	SaveUrlRepository func(*entity.Url) error
}

func Create(p CreateParameters) func(string) (*entity.Url, error) {
	return func(originalUrl string) (*entity.Url, error) {
		urlToken := p.CreateToken()
		var url = &entity.Url{
			Original:  originalUrl,
			Shortened: urlToken,
		}
		url.CreatedAt = time.Now()
		err := p.SaveUrlRepository(url)
		if err != nil {
			return nil, err
		}
		return url, nil
	}
}
