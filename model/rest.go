package model

// REST API structs

// GroupResponse defines a Group JSON response
type GroupResponse struct {
	Id         int    `json:"id"`
	Power      int    `json:"power"`
	Created    string `json:"created"`
	DeviceList []int  `json:"deviceList"`
}

// BlindResponse is the response from a blind GET.
type BlindResponse struct {
	DeviceMetadata DeviceMetadata `json:"deviceMetadata"`
	Position       float32        `json:"position"`
}

// RgbColorRequest allows (trying to) set a bulb color using classic hex RGB string.
type RgbColorRequest struct {
	RGBcolor string `json:"rgbcolor"`
}

// DimmingRequest allows setting the dimmer level from 0-255.
type DimmingRequest struct {
	Dimming int `json:"dimming"`
}

// PowerRequest contains a Power state int, 1 == on, 0 == off.
type PowerRequest struct {
	Power int `json:"power"`
}

// StateRequest allows setting both color, dimmer and power setting in a single PUT.
type StateRequest struct {
	RGBcolor string `json:"rgbcolor"`
	Dimmer   int    `json:"dimmer"`
	Power    int    `json:"power"`
}

// PositioningRequest allows setting the position from 0-100.
type PositioningRequest struct {
	Positioning float32 `json:"positioning"`
}
