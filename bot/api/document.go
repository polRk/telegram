// Bot API 5.7

package api

// Document is a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Thumb is a document thumbnail as defined by sender.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName is an original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}
