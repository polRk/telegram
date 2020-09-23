package telegram

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}
