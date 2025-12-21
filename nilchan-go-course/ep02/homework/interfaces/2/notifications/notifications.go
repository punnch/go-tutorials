// package notifications
// To improve iteraction with notifications
package notifications

// Notifier interface
// Stores notify methods
// (e.g. push, sms, email)
type Notifier interface {
	Send(to, message string)
}

// NotificationModule struct
// Stores a map
// - key is a name that user did for his message
// - value is a notify method of sending a message
type NotificationModule struct {
	notifiers map[string]Notifier
}

// NewNotifierModule()
// Used for encapsulation
// Returns a *NotificationModule with the created map field
func NewNotifierModule() *NotificationModule {
	return &NotificationModule{
		notifiers: make(map[string]Notifier),
	}
}

// AddNotifier()
// Adds a notifier in the map with an interface value
func (n *NotificationModule) AddNotifier(name string, notifier Notifier) {
	n.notifiers[name] = notifier
}

// SendAll()
// Recieves a notification map created by user
// Checks if it has a send method then uses a Send() function
func (n *NotificationModule) SendAll(notifications map[string]string) {
	for channel, message := range notifications {
		notifier, ok := n.notifiers[channel]
		if ok {
			notifier.Send("somebody", message)
		}
	}
}
