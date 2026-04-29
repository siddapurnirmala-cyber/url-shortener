package handler

import (
	"net/http"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	Service *service.URLService
}

func (h *URLHandler) Shorten(c *gin.Context) {

	var req struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	code, err := h.Service.CreateShortURL(req.URL)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"short_url": "http://localhost:9002/" + code,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {

	code := c.Param("code")

	url, err := h.Service.GetOriginalURL(code)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}
