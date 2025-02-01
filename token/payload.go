package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is unverifiable: error while executing keyfunc: token is invalid")
	ErrExpiredToken = errors.New("token has invalid claims: token is expired")
)

//Payload contains the payload data of the token

type Payload struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
	// ID        uuid.UUID `json:"id"`
	// IssuedAt  time.Time `json:"issued_at"`
	// ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		username,
		// ID:        tokenID,
		// Username:  username,
		// IssuedAt:  time.Now(),
		// ExpiredAt: time.Now().Add(duration),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        tokenID.String(),
		},
	}
	return payload, nil
}
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
