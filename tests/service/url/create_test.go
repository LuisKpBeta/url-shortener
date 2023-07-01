package test_service

import (
	"testing"
	"time"

	"github.com/LuisKpBeta/url-shortener/pk/entity"
	service_url "github.com/LuisKpBeta/url-shortener/pk/service/url"
)

func generateTokenMock(tokenValue string) func() string {
	return func() string {
		return tokenValue
	}
}
func generateRepositoryMock(id int) func(*entity.Url) error {
	return func(url *entity.Url) error {
		url.Id = id
		url.CreatedAt = time.Now()
		return nil
	}
}
func createSut(tokenValue string, id int) func(string) (*entity.Url, error) {
	tokenCreator := generateTokenMock(tokenValue)
	repository := generateRepositoryMock(id)
	create := service_url.Create(service_url.CreateParameters{
		CreateToken:       tokenCreator,
		SaveUrlRepository: repository,
	})
	return create
}

func TestCreateUrlService(t *testing.T) {
	originalUrl := "http://google.com"
	tokenToCreate := "fakeToken"
	create := createSut(tokenToCreate, 0)
	url, err := create(originalUrl)
	if err != nil {
		t.Fatalf("Create url fail with error: %s", err.Error())
	}
	if url.Original != originalUrl {
		t.Fatalf("Created url fail: wrong original url value")
	}
	if len(url.Shortened) < 1 || url.Shortened != tokenToCreate {
		t.Fatalf("Created url fail: invalid shortened url value")
	}
}
func TestCreateUrlService_ReturnIdFromRepository(t *testing.T) {
	originalUrl := "http://google.com"
	tokenToCreate := "fakeToken"
	id := 123
	create := createSut(tokenToCreate, id)
	url, err := create(originalUrl)
	if err != nil {
		t.Fatalf("Create url fail with error: %s", err.Error())
	}
	if url.Id != id {
		t.Fatalf("Created url fail: unexpected ID value")
	}
}
