// Bot API 5.7

package api

// User is a Telegram user or bot.
type User struct {
	// ID is a unique identifier for this user or bot.
	ID int64 `json:"id"`

	// IsBot is True, if use is a bot.
	IsBot bool `json:"is_bot"`

	// FirstName is a user's or bot's first name.
	FirstName string `json:"first_name"`

	// LastName is a user's or bot's last name.
	//
	// Optional.
	LastName string `json:"last_name,omitempty"`

	// Username is a user's or bot's username.
	//
	// Optional.
	Username string `json:"username,omitempty"`

	// LanguageCode is a IETF language tag of the user's language.
	// https://en.wikipedia.org/wiki/IETF_language_tag
	//
	// Optional.
	LanguageCode string `json:"language_code,omitempty"`

	// CanJoinGroups is True, if the bot can be invited to groups.
	// Returned only in api.GetMe().
	//
	// Optional.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages is True, if privacy mode is disabled for the bot.
	// Returned only in api.GetMe().
	//
	// Optional.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries is True, if the bot supports inline queries.
	// Returned only in api.GetMe().
	//
	// Optional.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}
