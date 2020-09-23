package telegram

type FoursquareType = string

type Venue struct {
	Location       *Location      `json:"location"`
	Title          string         `json:"title"`
	Address        string         `json:"address"`
	FoursquareID   string         `json:"foursquare_id,omitempty"`
	FoursquareType FoursquareType `json:"foursquare_type,omitempty"`
}
