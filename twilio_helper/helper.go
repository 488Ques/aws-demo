package twilio_helper

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

var client *twilio.RestClient

func TwillioInit() {
	accountSid := "AC50c8cfcb95e48227c6536e7a3febb559"
	authToken := "806287b8ff4d7ef7811074cbba807b07"

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func CreateMessage(message string) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetFrom("+12705801806")
	params.SetTo("+84347587031")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	response, _ := json.Marshal(*resp)
	fmt.Println("Response: " + string(response))
	return nil
}
