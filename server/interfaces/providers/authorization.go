package providers

type IAuthorizer interface {
	Authorize(accountID int64, token string) error
}

type AuthorizationProvider interface {
	GetAccountID(token string) (int64, error)
}
