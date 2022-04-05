// Bot API 5.7

package api

// File is a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
type File struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize in bytes, if known.
	//
	// Optional.
	FileSize int `json:"file_size,omitempty"`

	// FilePath, Use https://api.telegram.org/file/bot<token>/<file_path>.
	//
	// Optional.
	FilePath string `json:"file_path,omitempty"`
}
