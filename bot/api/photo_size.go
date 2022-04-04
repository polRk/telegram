// Bot API 5.7

package api

// PhotoSize is one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a photo width.
	Width int `json:"width"`

	// Height is a photo height.
	Height int `json:"height"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}
