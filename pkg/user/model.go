package user

// User domain model
type User struct {
	ID         int64
	UserID     string
	FirstName  string
	LastName   string
	Username   string
	Password   string
	Email      string
	IsDisabled bool
	IsVerified bool
}
