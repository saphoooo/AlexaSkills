package views

// These view represents Alexa custom-skills request type references
// https://developer.amazon.com/fr-FR/docs/alexa/custom-skills/request-types-reference.html#intentrequest

// AlexaRequest request from Alexa Skills
type AlexaRequest struct {
	Version string   `json:"version,omitempty"`
	Context *Context `json:"context,omitempty"`
	Session *Session `json:"session,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// Context ...
type Context struct {
	AudioPlayer *AudioPlayer `json:"AudioPlayer,omitempty"`
	System      *System      `json:"System,omitempty"`
}

// AudioPlayer ...
type AudioPlayer struct {
	PlayerActivity string `json:"playerActivity,omitempty"`
}

// System ...
type System struct {
	Application *Application `json:"application,omitempty"`
	User        *User        `json:"user,omitempty"`
	Device      *Device      `json:"device,omitempty"`
}

// Application ...
type Application struct {
	ApplicationID string `json:"applicationId"`
}

// User ...
type User struct {
	UserID string `json:"userId,omitempty"`
}

// Device ...
type Device struct {
	SupportedInterfaces *SupportedInterfaces `json:"supportedInterfaces,omitempty"`
}

// SupportedInterfaces ...
type SupportedInterfaces struct{}

// Session ...
type Session struct {
	New         bool         `json:"new,omitempty"`
	SessionID   string       `json:"sessionId,omitempty"`
	Application *Application `json:"application,omitempty"`
	User        *User        `json:"user,omitempty"`
	Attributes  struct{}     `json:"attributes,omitempty"`
}

// Request ...
type Request struct {
	Type        string  `json:"type,omitempty"`
	RequestID   string  `json:"requestId,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
	DialogState string  `json:"dialogState,omitempty"`
	Locale      string  `json:"locale,omitempty"`
	Intent      *Intent `json:"intent,omitempty"`
	Reason      string  `json:"reason,omitempty"`
	Error       *Error  `json:"error,omitempty"`
}

// Intent ...
type Intent struct {
	Name               string                 `json:"name,omitempty"`
	ConfirmationStatus string                 `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}

// SlotName ...
type SlotName struct {
	Name               string       `json:"name,omitempty"`
	Value              string       `json:"value,omitempty"`
	ConfirmationStatus string       `json:"confirmationStatus,omitempty"`
	Resolutions        *Resolutions `json:"resolutions,omitempty"`
}

// Resolutions ...
type Resolutions struct {
	ResolutionsPerAuthority []*ResolutionsPerAuthority `json:"resolutionsPerAuthority,omitempty"`
}

// ResolutionsPerAuthority ...
type ResolutionsPerAuthority struct {
	Authority string    `json:"authority,omitempty"`
	Status    *Status   `json:"status,omitempty"`
	Values    []*Values `json:"values,omitempty"`
}

// Status ...
type Status struct {
	Code string `json:"code,omitempty"`
}

// Values ...
type Values struct {
	Value *Value `json:"value,omitempty"`
}

// Value ...
type Value struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Error ...
type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

// Slot ...
type Slot struct {
	ConfirmationStatus string       `json:"confirmationStatus,omitempty"`
	Name               string       `json:"name,omitempty"`
	Resolutions        *Resolutions `json:"resolutions,omitempty"`
}
