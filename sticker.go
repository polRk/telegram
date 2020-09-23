package telegram

type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}
