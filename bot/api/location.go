// Bot API 5.7

package api

// Location is a point on the map.
type Location struct {
	// Longitude
	Longitude float64 `json:"longitude"`

	// Latitude
	Latitude float64 `json:"latitude"`

	// HorizontalAccuracy is the radius of uncertainty for the location, measured in meters.
	//
	// Optional.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod is the time in seconds relative to the message sending date, during which the location can be updated.
	// For active live locations only.
	//
	// Optional.
	LivePeriod int64 `json:"live_period,omitempty"`

	// Heading is the direction in degrees in which user is moving.
	// For active live locations only.
	//
	// Optional.
	Heading int64 `json:"heading,omitempty"`

	// ProximityAlertRadius is a maximum distance in meters for proximity alerts about approaching another chat member.
	//
	// Optional.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}

func (l *Location) Validate() bool {
	if l.HorizontalAccuracy < 0 || l.HorizontalAccuracy > 1500 {
		return false
	}

	if l.Heading < 0 || l.Heading > 360 {
		return false
	}

	return true
}
