package telegram

type MaskPositionPoint = string

const (
	MaskPositionPointForeHead MaskPositionPoint = "forehead"
	MaskPositionPointEyes     MaskPositionPoint = "eyes"
	MaskPositionPointMouth    MaskPositionPoint = "mouth"
	MaskPositionPointChin     MaskPositionPoint = "chin"
)

type MaskPosition struct {
	Point  MaskPositionPoint `json:"point"`
	XShift float64           `json:"x_shift"`
	YShift float64           `json:"y_shift"`
	Scale  float64           `json:"scale"`
}
