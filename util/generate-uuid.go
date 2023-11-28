package util

import "github.com/google/uuid"

func GenerateUUID() uuid.UUID {
	id, _ := uuid.NewRandom()
	return id
}
