package utils

import (
	"errors"

	"github.com/google/uuid"
)


func IsValidUUID(input string) bool {
	
	_, err := uuid.Parse(input)
	
	return errors.Is(err, nil)
}