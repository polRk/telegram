// Bot API 5.7

package api

type MaskPositionPoint = string

const (
	MaskPositionPointForeHead MaskPositionPoint = "forehead"
	MaskPositionPointEyes     MaskPositionPoint = "eyes"
	MaskPositionPointMouth    MaskPositionPoint = "mouth"
	MaskPositionPointChin     MaskPositionPoint = "chin"
)

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// Point is the part of the face relative to which the mask should be placed.
	// One of MaskPositionPointForeHead, MaskPositionPointEyes, MaskPositionPointMouth, or MaskPositionPointChin.
	Point MaskPositionPoint `json:"point"`

	// XShift is a shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float64 `json:"x_shift"`

	// YShift is a shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float64 `json:"y_shift"`

	// Scale is a mask scaling coefficient.
	// For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}
