package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/LuisKpBeta/url-shortener/internal/prometheus"
	"github.com/LuisKpBeta/url-shortener/pk/entity"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type CreateUrlSHortenerParams struct {
	Url string `json:"url" binding:"required"`
}

func CreateHttpServer() *gin.Engine {
	r := gin.Default()
	return r
}
func StartHttpServer(server *gin.Engine) {
	server.Use(gin.Logger())
	server.Run(":8080")
}
func CreateUrlShortnerHandler(server *gin.Engine, handler func(string) (*entity.Url, error)) {
	server.POST("/", func(c *gin.Context) {
		parameters := CreateUrlSHortenerParams{}
		err := c.Bind(&parameters)
		if err != nil {
			log.Println("[ERROR] on create shortener:" + err.Error())
			c.JSON(400, gin.H{
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
func GetUrlByTokenHandler(server *gin.Engine, handler func(token string) (string, error)) {
	server.GET("/:token", func(c *gin.Context) {
		token := c.Param("token")
		originalUrl, err := handler(token)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
			originalUrl = "http://" + originalUrl
		}

		c.Redirect(http.StatusMovedPermanently, originalUrl)
	})
}
func CreatePrometheusHandler(server *gin.Engine) {
	h := promhttp.Handler()
	server.GET("/metrics", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

}
func CreateMetricsMiddleware(server *gin.Engine, s *prometheus.PrometheusService) {
	server.Use(func(c *gin.Context) {
		metric := prometheus.NewHTTPMetric(c.Request.URL.Path, c.Request.Method)
		metric.Started()
		c.Next()
		metric.Finished()
		s.SaveHTTP(metric)
	})
}
