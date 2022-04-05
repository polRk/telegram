// Bot API 5.7

package api

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Text is an option text, 1-100 characters.
	Text string `json:"text"`

	// VoterCount is a number of users that voted for this option.
	VoterCount int `json:"voter_count"`
}

func (p PollOption) Validate() bool {
	if len(p.Text) < 1 || len(p.Text) > 100 {
		return false
	}

	return true
}
