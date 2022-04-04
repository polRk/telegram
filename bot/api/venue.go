// Bot API 5.7

package api

type Venue struct {
	// Location is a venue location.
	// Can't be a live location.
	Location *Location `json:"location"`

	// Title is a name of the venue.
	Title string `json:"title"`

	// Address of the venue.
	Address string `json:"address"`

	// Foursquare identifier of the venue.
	//
	// Optional.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType is type of the venue.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	//
	// Optional.
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID is a Google Places identifier of the venue.
	//
	// Optional.
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType is a Google Places type of the venue.
	// See: https://developers.google.com/maps/documentation/places/web-service/supported_types
	// Optional.
	GooglePlaceType string `json:"google_place_type,omitempty"`
}
