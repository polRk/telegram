// Bot API 5.7

package api

// PollAnswer is an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// PollID is a unique poll identifier.
	PollID string `json:"poll_id"`

	// User who changed the answer to the poll.
	User *User `json:"user"`

	// OptionIDs are 0-based identifiers of answer options, chosen by the user.
	// May be empty if the user retracted their vote.
	OptionIDs []int `json:"option_ids"`
}
