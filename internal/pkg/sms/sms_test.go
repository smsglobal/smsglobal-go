package sms

import (
	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/internal/util/mocks"
	"github.com/smsglobal/smsglobal-go/internal/util/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

var l *logger.Logger

func setup()  *client.Client{
	// Create the logger
	l = logger.CreateLogger(constants.DebugLevel)
	c := client.New("key", "secret")
	c.Logger = l
	l.Debug("Setup completed")

	return c
}

func TestSmsGetFailed(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	sms := &Client{
		Handler: c,
	}
	_, err := sms.Get("6746514019161950")

	assert.Error(t, err)
}

func TestSmsGetSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.SentToSingleDestinationResponse()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger: l,
	}

	res, err := sms.Get("6746514019161950")

	if err != nil {
		t.Errorf("Sms.Get returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.GetSmsResponse().Id, res.Id)
	assert.Equal(t, testdata.GetSmsResponse().OutgoingId, res.OutgoingId)
	assert.Equal(t, testdata.GetSmsResponse().Origin, res.Origin)
	assert.Equal(t, testdata.GetSmsResponse().Destination, res.Destination)
	assert.Equal(t, testdata.GetSmsResponse().Message, res.Message)
	assert.Equal(t, testdata.GetSmsResponse().Status, res.Status)
	assert.Equal(t, testdata.GetSmsResponse().DateTime, res.DateTime)
}

func TestSmsListFailed(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	sms := &Client{
		Handler: c,
		Logger: l,

	}
	_, err := sms.List(map[string]string{})

	assert.Error(t, err)
}

func TestSmsListSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.SmsListResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger: l,
	}

	options := make(map[string]string)
	options["status"] = "undelivered"
	options["destination"] = "61401869820"
	options["startDate"] = "2020-11-23 00:00:00"

	res, err := sms.List(options)

	if err != nil {
		t.Errorf("Sms.List returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.SmsListResponse().Total, res.Total)
	assert.Equal(t, testdata.SmsListResponse().Offset, res.Offset)
	assert.Equal(t, testdata.SmsListResponse().Limit, res.Limit)
	assert.Equal(t, testdata.SmsListResponse().Messages, res.Messages)
}

func TestSmsDeleteFailed(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNotFound,
	}

	sms := &Client{
		Handler: c,
		Logger: l,
	}
	err := sms.Delete("6746514019161950")
	assert.Error(t, err)
}

func TestSmsDeleteSuccess(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNoContent,
	}

	sms := &Client{
		Handler: c,
		Logger: l,
	}

	err := sms.Delete("6746514019161950")

	if err != nil {
		t.Errorf("Sms.Delete returned error: %v", err)
	}

	assert.Nil(t, err)
}

func TestSend(t *testing.T) {

}