package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserAuthTokenDetail struct {
	Raw       string    `json:"raw"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type UserAuthToken struct {
	TokenHash []byte    `json:"token"`
	UserId    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u UserAuthToken) IsExpired() bool {
	return u.ExpiresAt.Before(time.Now())
}

type UserAuthTokenCreate struct {
	TokenHash []byte    `json:"token"`
	UserId    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}
