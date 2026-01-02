package main

import (
	"study/notifications"
	"study/notifications/methods"
)

func main() {
	module := notifications.NewNotifierModule()

	// Add notifiers in a notifiers map
	module.AddNotifier("email", methods.NewEmail())
	module.AddNotifier("push", methods.NewPush())
	module.AddNotifier("sms", methods.NewSms())

	// Create the notification request
	notifications := map[string]string{
		"email": "You've got a million $!",
		"push":  "Your code is 8583, don't tell anyone!",
		"sms":   "Hey, don't forget to close buy a bread",
	}

	// Send the request
	module.SendAll(notifications)
}
