package models

import "github.com/google/uuid"

func generateUUID() string {
	return uuid.New().String()
}