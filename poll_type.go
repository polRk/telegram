package telegram

import "fmt"

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

// ErrBadPollTypeValue indicates an invalid value.
var ErrBadPollTypeValue = fmt.Errorf("telegram[PollType]: bad value")

type PollType string

// Validate returns an error if value is invalid.
func (p PollType) Validate() error {
	switch p {
	case PollTypeRegular, PollTypeQuiz:
		return nil
	}

	return ErrBadPollTypeValue
}
