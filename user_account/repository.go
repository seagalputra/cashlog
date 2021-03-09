package user_account

type UserAccountRepository interface {
	Save(account *UserAccount) error
	FindByID(id int64) (*UserAccount, error)
	FindByUsername(username string) (*UserAccount, error)
}
