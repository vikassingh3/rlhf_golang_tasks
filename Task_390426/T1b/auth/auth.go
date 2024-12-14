package auth

import (
	"my_project/user"
)

type Session struct {
	UserID string
}

func VerifyUser(userID string, session *Session) bool {
	return user.Verify(userID) // Use the user package
}
