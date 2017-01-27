package providers

import (
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"errors"
)

type Authorizer struct {
	authorizationProvider providers.AuthorizationProvider
}

func NewAuthorizer(authorizationProvider providers.AuthorizationProvider) *Authorizer {
	return &Authorizer{authorizationProvider: authorizationProvider}
}

func (authorizer *Authorizer) Authorize(accountID int64, token string) error {
	tokenAccountID, err := authorizer.authorizationProvider.GetAccountID(token)
	if err != nil {
		return err
	}
	if accountID != tokenAccountID {
		return errors.New("wrong account-id token")
	}
	return nil
}
