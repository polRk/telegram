package telegram

import "fmt"

const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML     ParseMode = "HTML"
)

// ErrBadAllowedUpdateValue indicates an invalid value.
var ErrBadParseModeValue = fmt.Errorf("telegram: bad value for ParseMode type")

// ParseMode represents how to parse entities in the message text.
type ParseMode string

// Validate returns an error if value is invalid.
func (value ParseMode) Validate() error {
	switch value {
	case ParseModeMarkdown, ParseModeHTML:
		return nil
	}

	return ErrBadParseModeValue
}
