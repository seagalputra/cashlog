package user_account

type UserAccountRepository interface {
	Save(account *UserAccount) error
	FindByID(id int64) (*UserAccount, error)
	FindByUsername(username string) (*UserAccount, error)
}

type UserAccountRepositoryImpl struct{}

func (r *UserAccountRepositoryImpl) FindByUsername(username string) (*UserAccount, error) {
	panic("implement me")
}

func (r *UserAccountRepositoryImpl) FindByID(id int64) (*UserAccount, error) {
	panic("implement me")
}

func (r *UserAccountRepositoryImpl) Save(account *UserAccount) error {
	panic("implement me")
}
