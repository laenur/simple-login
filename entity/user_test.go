package entity

import "testing"

func TestNewUser(t *testing.T) {
	username := "Admin"
	password := "Password"
	user := NewUser(username, password)

	if user.UserID != nil {
		t.Fatalf("expected user.UserID: nil, got: %d", *user.UserID)
	}
	if user.Username != username {
		t.Fatalf("expected user.Username: %s, got: %s", user.Username, username)
	}

	newUserWithSamePassword := NewUser(username, password)

	if newUserWithSamePassword.Salt == user.Salt {
		t.Fatalf("expected different salt, got: %s", newUserWithSamePassword.Salt)
	}
	if newUserWithSamePassword.Password == user.Password {
		t.Fatalf("expected different hashed password, got: %s", newUserWithSamePassword.Password)
	}
}
