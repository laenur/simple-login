package entity

import "testing"

func TestNewUser(t *testing.T) {
	username := "Admin"
	password := "Password"
	user := NewUser(username, password, []int{})

	if user.UserID != nil {
		t.Fatalf("expected user.UserID: nil, got: %d", *user.UserID)
	}
	if user.Username != username {
		t.Fatalf("expected user.Username: %s, got: %s", username, user.Username)
	}
	if user.Roles[0] != RoleUser {
		t.Fatalf("expected user.Roles: %d, got: %d", RoleUser, user.Roles[0])
	}

	newUserWithSamePassword := NewUser(username, password, []int{RoleAdmin})

	if newUserWithSamePassword.Salt == user.Salt {
		t.Fatalf("expected different salt, got: %s", newUserWithSamePassword.Salt)
	}
	if newUserWithSamePassword.Password == user.Password {
		t.Fatalf("expected different hashed password, got: %s", newUserWithSamePassword.Password)
	}
	if newUserWithSamePassword.Roles[0] != RoleAdmin {
		t.Fatalf("expected newUserWithSamePassword.Roles: %d, got: %d", RoleAdmin, newUserWithSamePassword.Roles[0])
	}
}
