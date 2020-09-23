package telegram

type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}
