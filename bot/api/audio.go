// Bot API 5.7

package api

// Audio is an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`

	// Performer of the audio as defined by sender or by audio tags.
	//
	// Optional.
	Performer string `json:"performer,omitempty"`

	// Title of the audio as defined by sender or by audio tags.
	//
	// Optional.
	Title string `json:"title,omitempty"`

	// FileName is an original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is MIME type of the file as defined by sender
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`

	// Thumb is a thumbnail of the album cover to which the music file belongs.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}
