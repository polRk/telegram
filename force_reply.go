package telegram

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// Validate returns an error if value is invalid.
func (k *ForceReply) Validate() error {
	return nil
}
