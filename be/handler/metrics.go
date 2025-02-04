package handler

import (
	"fmt"
	"net/http"
	"time"
	"udp-be/db"
	"udp-be/models"

	"github.com/gin-gonic/gin"
)

func HandleMetrics(ctx *gin.Context) {
	var data []models.MetricsPayload
	tenMinutesAgo := time.Now().Add(-10 * time.Minute)
	result := db.DB.Table("message").Where("created_at >= ?", tenMinutesAgo).Find(&data)
	if result.Error != nil {
		fmt.Println("Failed reading from DB", result.Error)
	}
	ctx.JSON(http.StatusOK, data)
}
