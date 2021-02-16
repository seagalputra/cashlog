package user_account

type UserAccountRepository interface {
	InsertUserAccount(account *UserAccount)
}

type UserAccountRepositoryImpl struct{}

func (r *UserAccountRepositoryImpl) InsertUserAccount(account *UserAccount) {
	panic("implement me")
}
