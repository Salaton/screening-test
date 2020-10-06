package notification

import (
	"fmt"

	"github.com/davidmutia47/AfricasTalkingGateway"
)

func SendNotification(customername, phonenumber string) string {
	// var DB db.DBClient
	// DB.FetchPhoneNumber()
	// Specify your login credentials
	username := "drowsydriver"
	apikey := "a7f07c892fb5a74e5b90481a10a72154893f216fd85c71e2271b21167fa56556"
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
