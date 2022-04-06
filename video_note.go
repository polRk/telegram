// Bot API 5.7

package telegram

// VideoNote is a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Length is a video width and height (diameter of the video message) as defined by sender.
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb is a video thumbnail.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}
