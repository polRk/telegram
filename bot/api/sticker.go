// Bot API 5.7

package api

// Sticker is a sticker.
type Sticker struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is a sticker width.
	Width int `json:"width"`

	// Height is a sticker height.
	Height int `json:"height"`

	// IsAnimated is True, if the sticker is animated.
	IsAnimated bool `json:"is_animated"`

	// IsVideo is True, if the sticker is video sticker.
	IsVideo bool `json:"is_video"`

	// Thumb is a sticker thumbnail in the .WEBP or .JPG format.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Emoji associated with the sticker.
	//
	// Optional.
	Emoji string `json:"emoji,omitempty"`

	// SetName is a name of the sticker set to which the sticker belongs.
	//
	// Optional.
	SetName string `json:"set_name,omitempty"`

	// MaskPosition is the position where the mask should be placed.
	// For mask stickers.
	//
	// Optional.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// FileSize in bytes.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`
}
