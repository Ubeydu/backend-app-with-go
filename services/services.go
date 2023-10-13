package services

import (
    "github.com/Ubeydu/RESTAppWithGo/models"
)

var storage []models.Item = []models.Item{}

// GetAllItems returns all items
func GetAllItems() []models.Item {
	return storage
}
