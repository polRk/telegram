package telegram

// WebhookInfo represents information about the current status of a webhook.
type WebhookInfo struct {
	URL                  string          `json:"url"`
	HasCustomCertificate bool            `json:"has_custom_certificate,omitempty"`
	PendingUpdateCount   int             `json:"pending_update_count,omitempty"`
	LastErrorDate        int             `json:"last_error_date,omitempty"`
	LastErrorMessage     int             `json:"last_error_message,omitempty"`
	MaxConnections       int             `json:"max_connections,omitempty"`
	AllowedUpdates       []AllowedUpdate `json:"allowed_updates,omitempty"`
}
