package model

// General defines (with JSON tags) general attributes of IKEA trådfri entity
type General struct {
	Name      string `json:"9001"`
	CreatedAt int    `json:"9002"`
	DeviceId  int    `json:"9003"`
}

// Metadata defines (with JSON tags) metadata attributes of IKEA trådfri entity
type Metadata struct {
	Vendor   string `json:"0"`
	TypeName string `json:"1"`
	Num2     string `json:"2"`
	TypeID   string `json:"3"`
	Num6     int    `json:"6"`
	Battery  int    `json:"9"`
}

// LightControl defines (with JSON tags) a IKEA trådfri LightControl.
type LightControl struct {
	ID         int    `json:"9003"`
	RGBHex     string `json:"5706"`
	Hue        int    `json:"5707"`
	Saturation int    `json:"5708"`
	CIE_1931_X int    `json:"5709"`
	CIE_1931_Y int    `json:"5710"`
	Power      int    `json:"5850"`
	Dimmer     int    `json:"5851"`
}

// BlindControl defines (with JSON tags) a IKEA trådfri BlindControl.
type BlindControl struct {
	Position float32 `json:"5536"`
	ID       int     `json:"9003"`
}

// Device defines (with JSON tags) a IKEA trådfri device of some kind
type Device struct {
	General

	Metadata     Metadata       `json:"3"`
	LightControl []LightControl `json:"3311"`
	BlindControl []BlindControl `json:"15015"`

	ApplicationType int `json:"5750"`
	ReachableState  int `json:"9019"`
	LastSeen        int `json:"9020"`
	Num9054         int `json:"9054"`
}

// Group defines (with JSON tags) a IKEA trådfri Group.
type Group struct {
	General

	Power   int `json:"5850"`
	Num5851 int `json:"5851"`

	Content struct {
		DeviceList struct {
			DeviceIds []int `json:"9003"`
		} `json:"15002"`
	} `json:"9018"`

	Num9039 int `json:"9039"`
	Num9108 int `json:"9108"`
}

// RemoteControl defines (with JSON tags) a IKEA remote control.
type RemoteControl struct {
	General

	Metadata Metadata `json:"3"`

	ApplicationType int `json:"5750"`
	ReachableState  int `json:"9019"`
	LastSeen        int `json:"9020"`
	Num9054         int `json:"9054"`
	Num15009        []struct {
		ID int `json:"9003"`
	} `json:"15009"`
}

// ControlOutlet defines (with JSON tags) a IKEA control outlet.
type ControlOutlet struct {
	General

	Metadata     Metadata `json:"3"`
	PowerControl []struct {
		Num5850 int `json:"5850"`
		Num5851 int `json:"5851"`
		ID      int `json:"9003"`
	} `json:"3312"`

	ApplicationType int    `json:"5750"`
	ReachableState  int    `json:"9019"`
	LastSeen        int    `json:"9020"`
	Num9054         int    `json:"9054"`
	Num9084         string `json:"9084"`
}

// DeviceMetadata defines (with JSON tags) common device metadata. Typically embedded in other structs.
type DeviceMetadata struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Type    string `json:"type"`
	Battery int    `json:"battery"`
}

// PowerPlugResponse is the response from a power plug device GET.
type PowerPlugResponse struct {
	DeviceMetadata DeviceMetadata `json:"deviceMetadata"`
	Power          bool           `json:"power"`
}

// BulbResponse is the response from a light bulb GET.
type BulbResponse struct {
	DeviceMetadata DeviceMetadata `json:"deviceMetadata"`
	Dimmer         int            `json:"dimmer"`
	CIE_1931_X     int            `json:"xcolor"`
	CIE_1931_Y     int            `json:"ycolor"`
	RGB            string         `json:"rgbcolor"`
	Power          bool           `json:"power"`
}

// Result is a generic result containing a plain text message
type Result struct {
	Msg string
}

// TokenExchange maps the human-readable Token and TypeIdentifies into their IKEA specific numeric codes.
type TokenExchange struct {
	Token          string `json:"9091"`
	TypeIdentifier string `json:"9029"`
}
