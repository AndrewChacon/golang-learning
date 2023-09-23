package main

import (
	"fmt"
)
 
func main () {

	// Variables and Conditionals

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

	// Functions

	// func sub(x int, y int) int {
	// 	return x - y
	// }

	
	// fmt.Println(concat("Hello", "World"))

	// fmt.Println(add(1,2))

	// sendsSoFar := 430
	// const sendsToAdd = 25
	// sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	// fmt.Println("You've sent", sendsSoFar, "messages")

	// firstName, _, _ := getNames()
	// fmt.Println("Welcome to Textio", firstName)

	// STRUCTS

	// type car struct {
	// 	Make string
	// 	Model string
	// 	Height int
	// 	width int
	// }

	// type messageToSend struct {
	// 	phoneNumber int
	// 	message string
	// }

	// type Wheel struct {
	// 	Radius int
	// 	Material string
	// }

	// type Car struct {
	// 	Make string
	// 	Model string
	// 	Height int
	// 	Width int
	// 	Frontwheel Wheel
	// 	Backwheel Wheel
	// }

	// myCar := Car{}
	// myCar.Frontwheel.Radius = 5

	// type Car struct { 
	// 	Make string
	// 	Model string
	// 	Height int
	// 	Width int
	// 	Wheel struct {
	// 		Radius int
	// 		Material string
	// 	}
	// }

	type car struct {
		make string
		model string
	}
	
	type truck struct {
		car
		bedSize int
	}

	myTruck := truck{}
	myTruck.make = "a"
	
}

type User struct {
	name string
	number int
}

type messageToSend struct {
	message string
	sender User 
	recipient User
}

func canSendMessage(mToSend messageToSend) bool {
	if(mToSend.sender.name == "" && mToSend.sender.number == 0) {
		return false
	}
	if(mToSend.recipient.name == "" && mToSend.recipient.number == 0) {
		return false
	}
	return false
}

// func divide(dividen, divisor int) (int, error) {
// 	if divisor == 0 {
// 		return 0, errors.New("Cant divide by 0") // returns early with an error
// 	}
// 	return dividen/divisor, nil // nil error means no error 
// }

// func getCoords()(x, y int) {
// 	return // naked return, automatically returns x and y
// }

// func yearsUntilEvents(age int)(yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
// 	yearsUntilAdult = 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}

// 	yearsUntilDrinking = 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}

// 	yearsUntilCarRental = 25 - age
// 	if yearsUntilCarRental < 0 {
// 		yearsUntilCarRental = 0
// 	}

// 	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
// }

// func getNames()(string, string, string) { // multiple return values
// 	return "Andrew", "Bolivar", "Chacon"
// }

// func incrementSends(sendsSoFar, sendsToAdd int) int {
// 	return sendsSoFar + sendsToAdd
// }

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
	
// }

// func add(x, y int) int {
// 	return x + y
// }

// Data Types

// bool
// string 
// int int8 int16 int32 int64
// float32 float 64
