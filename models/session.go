package models

import (
	"database/sql"
	"fmt"

	"github.com/JozeFons/Web_Development_with_Go_v2/rand"
)

// The minimum number of bytes to be used for each session token.
const (
	MinBytesPerToken = 32
)

// Token is only set when creating a new session. When look up a session
// this will be left empty, as we only store the hash of a session token
// in our database and we cannot reverse it into a raw token.
type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID: userID,
		Token:  token,
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
