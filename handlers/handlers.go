package handlers

import (
	"net/url"
	"url_shortener/models"
	"url_shortener/repository"
	"url_shortener/shortcode"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	repo *repository.URLRepository
}

func NewUrlHandler(repo *repository.URLRepository) *UrlHandler {
	return &UrlHandler{repo: repo}
}

func (h *UrlHandler) CreateShortURL(c *gin.Context) {
    var request struct {
        URL string `json:"url"`
    }
    
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
    
    if _, err := url.ParseRequestURI(request.URL); err != nil {
        c.JSON(400, gin.H{"error": "Invalid URL"})
        return
    }
    
    shortCode := shortcode.GenerateUnique(func(code string) bool {
		exists, _ := h.repo.Exists(code)
		return exists
	})
    
    newURL := models.Url{
        OriginalURL: request.URL,
        ShortCode:   shortCode,
    }
    
    if err := h.repo.Create(&newURL); err != nil {
        c.JSON(500, gin.H{"error": "Failed to create short URL"})
        return
    }
    
    c.JSON(201, newURL)
}

func (h *UrlHandler) GetOriginalURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	
	url, err := h.repo.FindByShortCode(shortCode)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get original URL"})
		return
	}
	
	if url == nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	
	c.JSON(200, url)
}

func (h *UrlHandler) UpdateShortURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	
	var request struct {
		URL string `json:"url"`
	}
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	
	if _, err := url.ParseRequestURI(request.URL); err != nil {
		c.JSON(400, gin.H{"error": "Invalid URL"})
		return
	}
	
	url, err := h.repo.FindByShortCode(shortCode)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update short URL"})
		return
	}
	
	if url == nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	
	url.OriginalURL = request.URL
	if err := h.repo.Update(url); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update short URL"})
		return
	}
	
	c.JSON(200, url)
}

func (h *UrlHandler) DeleteShortURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	
	url, err := h.repo.FindByShortCode(shortCode)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete short URL"})
		return
	}
	
	if url == nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	
	if err := h.repo.Delete(url); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete short URL"})
		return
	}
	
	c.JSON(204, nil)
}

func (h *UrlHandler) GetStats(c *gin.Context) {
	shortCode := c.Param("shortCode")
	
	url, err := h.repo.FindByShortCode(shortCode)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get stats"})
		return
	}
	
	if url == nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	
	c.JSON(200, gin.H{
		"accessCount": url.AccessCount,
		"createdAt":   url.CreatedAt,
		"updatedAt":   url.UpdatedAt,
	})
}