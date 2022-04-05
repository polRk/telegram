// Bot API 5.7

package api

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

type PollType string

// Poll contains information about a poll.
type Poll struct {
	// ID is a unique poll identifier.
	ID string `json:"id"`

	// Question is a Poll question, 1-300 characters.
	Question string `json:"question"`

	// Options is list of poll options.
	Options []*PollOption `json:"options"`

	// TotalVoterCount is a total number of users that voted in the poll.
	TotalVoterCount int `json:"total_voter_count"`

	// IsClosed is True, if the poll is closed.
	IsClosed bool `json:"is_closed"`

	// IsAnonymous is True, if the poll is anonymous.
	IsAnonymous bool `json:"is_anonymous"`

	// Type is a PollType, currently can be PollTypeRegular or PollTypeQuiz
	Type PollType `json:"type"`

	// AllowsMultipleAnswers is True, if the poll allows multiple answers.
	//
	// Optional.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// CorrectOptionID is a 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode,
	// which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	//
	// Optional.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Explanation is a text that is shown when a user chooses an incorrect answer or taps on the lamp icon
	// in a quiz-style poll, 0-200 characters.
	//
	// Optional.
	Explanation string `json:"explanation,omitempty"`

	// ExplanationEntities is a special entities like usernames, URLs, bot commands, etc. that appear in the explanation.
	//
	// Optional.
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`

	// OpenPeriod is an amount of time in seconds the poll will be active after creation.
	//
	// Optional.
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate is a point in time (Unix timestamp) when the poll will be automatically closed.
	//
	// Optional.
	CloseDate int `json:"close_date,omitempty"`
}

func (p *Poll) Validate() bool {
	if len(p.Question) < 1 || len(p.Question) > 300 {
		return false
	}

	if len(p.Explanation) > 200 {
		return false
	}

	return true
}
