package user_account

type UserAccount struct {
	Id         int64
	UserId     string
	FirstName  string
	LastName   string
	Username   string
	Password   string
	Email      string
	IsDisabled bool
	IsVerified bool
}
