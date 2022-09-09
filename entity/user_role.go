package entity

const (
	RoleAdmin int = 1000
	RoleUser  int = 1
)

type UserRole struct {
	UserRoleID int64
	UserID     int64
	RoleID     int
}
