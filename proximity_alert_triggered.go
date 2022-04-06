// Bot API 5.7

package telegram

// ProximityAlertTriggered is the content of a service message,
// sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// Traveler is the User that triggered the alert.
	Traveler *User `json:"traveler"`

	// Watcher is the User that set the alert.
	Watcher *User `json:"watcher"`

	// Distance between the users.
	Distance int `json:"distance"`
}
