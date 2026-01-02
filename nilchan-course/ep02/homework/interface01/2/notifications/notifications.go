package notifications

// Notifier interface stores notify methods (e.g. push, sms, email)
type Notifier interface {
	Send(to, message string)
}

// NotificationModule struct stores a map for different channels
type NotificationModule struct {
	notifiers map[string]Notifier
}

// NewNotifierModul creates a new NotificationModule with an empty Notifier map
func NewNotifierModule() *NotificationModule {
	return &NotificationModule{
		notifiers: make(map[string]Notifier),
	}
}

// AddNotifier adds a Notifier with the given name to the map
func (n *NotificationModule) AddNotifier(name string, notifier Notifier) {
	n.notifiers[name] = notifier
}

// SendAll sends all the messages of a recieved map
func (n *NotificationModule) SendAll(notifications map[string]string) {
	for channel, message := range notifications {
		notifier, ok := n.notifiers[channel]
		if ok {
			notifier.Send("somebody", message)
		}
	}
}
