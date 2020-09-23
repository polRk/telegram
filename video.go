package telegram

type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}
