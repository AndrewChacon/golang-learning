package main

import (
	"fmt"
)
 
func main () {
	// var smsSendingLimit int
	// var costPerSMS float64
	// var hasPermission bool
	// var username string

	// fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// empty := ""
	// fmt.Println(empty)

	averageOpenRate, displayMessage := 0.23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)

	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years.")

	const premiumPlanName = "Preimium Plan"
	const basicPlanName = "Basic Plan"
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour
	fmt.Println("Number of seconds in an hour:", secondsInHour)

	const name = "Andrew Chacon"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
	fmt.Println(msg)

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

// Data Types

// bool
// string 
// int int8 int16 int32 int64
// float32 float 64
