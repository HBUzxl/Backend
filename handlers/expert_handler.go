package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExpertsHandler(c *gin.Context) {
	experts, err := services.GetExperts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Experts " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"experts": experts,
	})
}
