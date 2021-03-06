package user

import (
	"net/http"

	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
)

var path = "/user/credit-balance"

type Client struct {
	Handler *client.Client
	Logger  *logger.Logger
}

// CreditBalance method performs API request to get a user account balance
func (c *Client) CreditBalance() (*api.BalanceResponse, error) {

	log := c.Logger.Lgr.With().Str("USER API Layer", "CreditBalance").Logger()

	log.Debug().Msg("Initiating account balance request")

	req, err := c.Handler.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	balance := &api.BalanceResponse{}
	err = c.Handler.Do(req, balance)

	if err != nil {
		return nil, err
	}

	return balance, nil
}
