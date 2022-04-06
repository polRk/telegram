// Bot API 5.7

package telegram

// UserProfilePhotos is a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount is a total number of profile pictures the target user has.
	TotalCount int `json:"total_count"`

	// Photos is an array of requested profile pictures (in up to 4 sizes each)
	Photos []*PhotoSize `json:"photos"`
}
