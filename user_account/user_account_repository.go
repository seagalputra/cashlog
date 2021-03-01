package user_account

type UserAccountRepository interface {
	Save(account *UserAccount) error
}

type UserAccountRepositoryImpl struct{}

func (r *UserAccountRepositoryImpl) Save(account *UserAccount) error {
	panic("implement me")
}
