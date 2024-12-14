package session

import (
	"my_project/user"
)

type Session struct {
	UserID string
}

func NewSession(userID string) *Session {
	if !user.Verify(userID) {
		return nil // Invalid user
	}
	return &Session{UserID: userID}
}
