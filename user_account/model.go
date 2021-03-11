package user_account

type UserAccount struct {
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
