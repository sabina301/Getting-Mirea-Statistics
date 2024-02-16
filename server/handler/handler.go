package handler

import (
	"getting-statistics-mirea/entity"
	"getting-statistics-mirea/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) GetResult(c *gin.Context) {

	if err := c.BindJSON(&entity.User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	statistics := h.service.GetResultService(entity.User.Email, entity.User.Password, entity.User.Number)
	c.JSON(http.StatusOK, statistics)
}

func (h *Handler) GetPage(c *gin.Context) {
	htmlBytes, err := ioutil.ReadFile("../client/index.html")

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	htmlContent := string(htmlBytes)
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, htmlContent)
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
