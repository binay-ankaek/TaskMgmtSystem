package app

import (
	"task-service/internal/domain"
)

func extractUserIds(users []domain.UserDetails) []string {
	// loop among users and extract user.ID.
	var id []string
	for _, user := range users {
		id = append(id, user.ID)

	}
	return id
}
