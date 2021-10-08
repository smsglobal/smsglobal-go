package smsglobal

import (
	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/pkg/sms"
	"github.com/smsglobal/smsglobal-go/internal/pkg/user"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	e "github.com/smsglobal/smsglobal-go/pkg/error"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
)

// SMSGlobal defines the SMSGlobal client.
type SMSGlobal struct {
	User *user.Client
	Sms  *sms.Client
}

// New Init initializes the SMSGlobal client with all available resources
func New(key, secret string) (*SMSGlobal, error) {

	// Create the logger
	l := logger.CreateLogger(constants.DebugLevel)
	lg := l.Lgr.With().Str("SMSGlobal", "New").Logger()
	lg.Info().Msgf("Creating SMSGlobal instance")

	if key == "" || secret == "" {
		return nil, &e.Error{Message: "API key and Secret are required!", Code: constants.DefaultCode}
	}

	s := new(SMSGlobal)

	c := client.New(key, secret)
	c.Logger = l
	s.User = &user.Client{Handler: c, Logger: l}

	s.Sms = &sms.Client{Handler: c, Logger: l}

	return s, nil
}

// CreateSms Creates an empty Sms object. Populate relevant properties for sending a message
func (s *SMSGlobal) CreateSms() *api.SendSingleSms {
	return &api.SendSingleSms{}
}

// CreateMultipleSms Creates an empty MultipleSms object. Populated relevant properties for sending a message
func (s *SMSGlobal) CreateMultipleSms() *api.SendMultipleSms {
	return &api.SendMultipleSms{}
}
