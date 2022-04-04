// Bot API 5.7

package api

// Video is a video file.
type Video struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a video width as defined by sender.
	Width int `json:"width"`

	// Height is a video height as defined by sender.
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb is a video thumbnail.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName is an original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name,omitempty"`

	// MimeType is Mime type of file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}
