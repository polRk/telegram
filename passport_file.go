// Bot API 5.7

package telegram

// PassportFile is a file uploaded to Telegram Passport.
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize in bytes.
	FileSize int `json:"file_size"`

	// FileDate is a unix time when the file was uploaded.
	FileDate int `json:"file_date"`
}
