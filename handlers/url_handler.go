package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiansibatuara/golang-url-shortener/service"
)

type UrlHandler struct {
	UrlService service.UrlService
}

func NewUrlHandler(urlService service.UrlService) *UrlHandler {
	return &UrlHandler{
		UrlService: urlService,
	}
}

func (h *UrlHandler) CreateShortUrl(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	newURL, err := h.UrlService.CreateUrl(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newURL)
}

func (h *UrlHandler) RedirectToOriginal(c *gin.Context) {
	code := c.Param("code")
	url, err := h.UrlService.GetOriginalUrl(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.Url)
}

func (h *UrlHandler) GetStats(c *gin.Context) {
	code := c.Param("code")
	stats, err := h.UrlService.GetStats(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *UrlHandler) UpdateShortUrl(c *gin.Context) {
	code := c.Param("code")
	var req struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	newURL, err := h.UrlService.UpdateUrl(code, req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newURL)
}

func (h *UrlHandler) DeleteShortUrl(c *gin.Context) {
	code := c.Param("code")
	if err := h.UrlService.DeleteUrl(code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
