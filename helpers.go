package main

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func InitializeTwillio() *twilio.RestClient {
	accountSid := "AC50c8cfcb95e48227c6536e7a3febb559"
	authToken := "85061e2f857d04779fb834bac7074660"

	return twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func CreateMessage(message string, client *twilio.RestClient) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetFrom("+12705801806")
	params.SetTo("+84347587031")
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
