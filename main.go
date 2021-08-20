package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	ToPhoneNumber := "+94778435151"
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	client := twilio.NewRestClient()
	params := &openapi.CreateMessageParams{}
	params.SetTo(ToPhoneNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("Hi , sending from twilio test!")

	_, err = client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}