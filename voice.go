// Bot API 5.7

package telegram

// Voice is a voice note.
type Voice struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`

	// MimeType is MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}
