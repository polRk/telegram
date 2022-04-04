// Bot API 5.7

package api

// ChatLocation is a location to which a chat is connected.
type ChatLocation struct {
	// Location is the location to which the supergroup is connected.
	// Can't be a live location.
	Location *Location

	// Address is the location address, defined by the chat owner.
	Address string `json:"address"`
}

func (l *ChatLocation) Validate() bool {
	if len(l.Address) < 1 || len(l.Address) > 64 {
		return false
	}

	return true
}
