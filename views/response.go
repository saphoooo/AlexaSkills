package views

// These view represents Alexa custom-skills request type references
// https://developer.amazon.com/fr-FR/docs/alexa/custom-skills/request-types-reference.html#intentrequest

// AlexaResponse ...
type AlexaResponse struct {
	Version           string             `json:"version"`
	SessionAttributes *SessionAttributes `json:"sessionAttributes,omitempty"`
	Response          *Response          `json:"response"`
}

// SessionAttributes ...
type SessionAttributes struct {
	Key string `json:"key,omitempty"`
}

// Response ...
type Response struct {
	ShouldEndSession bool          `json:"shouldEndSession,omitempty"`
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []*Directives `json:"directives,omitempty"`
}

// OutputSpeech ...
type OutputSpeech struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

// Card ...
type Card struct {
	Type  string `json:"type"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Image *Image `json:"image,omitempty"`
}

// Image ...
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// Reprompt ...
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

// Directives ...
type Directives struct {
	Type string `json:"type,omitempty"`
}
