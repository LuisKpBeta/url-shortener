package api

import (
	"log"

	"github.com/LuisKpBeta/url-shortener/pk/entity"
	"github.com/gin-gonic/gin"
)

type CreateUrlSHortenerParams struct {
	Url string `json:"url" binding:"required"`
}

func CreateHttpServer() *gin.Engine {
	r := gin.Default()
	return r
}
func StartHttpServer(server *gin.Engine) {
	server.Run(":8080")
}
func CreateUrlShortnerHandler(server *gin.Engine, handler func(string) (*entity.Url, error)) {
	server.POST("/", func(c *gin.Context) {
		parameters := CreateUrlSHortenerParams{}
		c.Bind(&parameters)
		err := c.Bind(&parameters)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		entity, err := handler(parameters.Url)
		if err != nil {
			log.Println("[ERROR] on create shortener:" + err.Error())
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			return
		}
		c.JSON(200, entity)
	})
}
