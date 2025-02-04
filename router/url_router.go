package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tiansibatuara/golang-url-shortener/handlers"
)

func NewUrlRouter(r *gin.Engine, urlController *handlers.UrlHandler) {
	r.POST("/shorten", urlController.CreateShortUrl)
	r.GET("/:code", urlController.RedirectToOriginal)
	r.GET("/:code/stats", urlController.GetStats)
	r.PUT("/:code", urlController.UpdateShortUrl)
	r.DELETE("/shorten/:code", urlController.DeleteShortUrl)
}
