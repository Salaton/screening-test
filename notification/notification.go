package notification

import (
	"fmt"
	"os"

	"github.com/davidmutia47/AfricasTalkingGateway"
)

func SendNotification(customername, phonenumber string) string {
	// var DB db.DBClient
	// DB.FetchPhoneNumber()
	// Specify your login credentials
	username := os.Getenv("AFRICASTALKINGUSERNAME")
	apikey := os.Getenv("AFRICASTALKINGAPIKEY")
	// Specify the numbers that you want to send to in a comma-separated list
	// Please ensure you include the country code (+254 for Kenya in this case)
	recipients := phonenumber
	// And of course we want our recipients to know what we really do
	message := "Hello there " + customername + " Your order has been received"

	//Create instance of getWay
	getWay := AfricasTalkingGateway.AfricasTalkingGateway(username, apikey)

	//sandbox
	//getWay := AfricasTalkingGateway.AfricasTalkingGateway(username,apikey,"sandbox")

	//call sendMessage to handle sending the message
	response, err := getWay.SendMessage(recipients, message)

	//handle errors if encountered an error
	if err != nil {
		//handle error
	}
	fmt.Println(response)
	return message
}
