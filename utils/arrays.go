package utils

import (
	"golang.org/x/exp/constraints"
	"terminal_commands/models"
)

type Type interface {
	constraints.Ordered | models.User | models.Command | models.UserResponse | models.CommandResponse | models.Platform
}

func Map[T Type, V Type](values []T, fn func(T) V) []V {
	result := make([]V, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}
