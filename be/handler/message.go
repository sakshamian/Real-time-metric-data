package handler

import (
	"fmt"
	"udp-be/db"
	"udp-be/models"
)

func WriteToDatabase(data models.MetricsPayload) {
	result := db.DB.Table("message").Create(&data)
	if result.Error != nil {
		fmt.Println("Failed writing to DB:", result.Error)
		return
	}
	fmt.Println("Successfully saved to DB:", data)
}
