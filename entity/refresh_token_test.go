package entity

import (
	"testing"
	"time"
)

func TestNewRefreshToken(t *testing.T) {
	userID := int64(1)
	refreshToken := NewRefreshToken(userID)

	if refreshToken.UserID != userID {
		t.Fatalf("expected refreshToken.UserID: %d, got: %d", userID, refreshToken.UserID)
	}

	timeIn6Days := time.Now().Add(time.Hour * 24 * 6)
	if !refreshToken.ValidUntil.After(timeIn6Days) {
		t.Fatal("expected refreshToken.ValidUntil more than 6 day")
	}

	newRefreshToken := NewRefreshToken(userID)
	if refreshToken.Token == newRefreshToken.Token {
		t.Fatal("expected different token for newRefreshToken")
	}
}
