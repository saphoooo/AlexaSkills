package google

// AcionsRequest ...
type AcionsRequest struct {
	ResponseID                  string                       `json:"responseId,omitempty"`
	QueryResult                 *QueryResult                 `json:"queryResult,omitempty"`
	OriginalDetectIntentRequest *OriginalDetectIntentRequest `json:"originalDetectIntentRequest,omitempty"`
	Session                     string                       `json:"session,omitempty"`
}

// QueryResult ...
type QueryResult struct {
	QueryText                 string                 `json:"queryText,omitempty"`
	Action                    string                 `json:"action,omitempty"`
	Parameters                *Parameters            `json:"parameters,omitempty"`
	AllRequiredParamsPresent  bool                   `json:"allRequiredParamsPresent,omitempty"`
	FulfillmentMessages       []*FulfillmentMessages `json:"fulfillmentMessages,omitempty"`
	OutputContexts            []*OutputContexts      `json:"outputContexts,omitempty"`
	Intent                    *Intent                `json:"intent,omitempty"`
	IntentDetectionConfidence float32                `json:"intentDetectionConfidence,omitempty"`
	LanguageCode              string                 `json:"languageCode,omitempty"`
}

// Parameters ...
// TODO should be improuved with an interface
type Parameters struct {
	Foods             string `json:"Foods,omitempty"`
	DietTypes         string `json:"DietTypes,omitempty"`
	FoodsOriginal     string `json:"Foods.original,omitempty"`
	DietTypesOriginal string `json:"DietTypes.original,omitempty"`
}

// FulfillmentMessages ...
type FulfillmentMessages struct {
	Text struct {
		Text []string `json:"text,omitempty"`
	} `json:"text,omitempty"`
}

// OutputContexts ...
type OutputContexts struct {
	Name       string      `json:"name,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

// Intent ...
type Intent struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// OriginalDetectIntentRequest ...
type OriginalDetectIntentRequest struct {
	Source  string   `json:"source,omitempty"`
	Version string   `json:"version,omitempty"`
	Payload *Payload `json:"payload,omitempty"`
}

// Payload ...
type Payload struct {
	User              *User                `json:"user,omitempty"`
	Conversation      *Conversation        `json:"conversation,omitempty"`
	Inputs            []*Inputs            `json:"inputs,omitempty"`
	Surface           *Surface             `json:"surface,omitempty"`
	IsInSandbox       bool                 `json:"isInSandbox,omitempty"`
	AvailableSurfaces []*AvailableSurfaces `json:"availableSurfaces,omitempty"`
	RequestType       string               `json:"requestType,omitempty"`
}

// User ...
type User struct {
	Locale                 string `json:"locale,omitempty"`
	LastSeen               string `json:"lastSeen,omitempty"`
	UserVerificationStatus string `json:"userVerificationStatus,omitempty"`
}

// Conversation ...
type Conversation struct {
	ConversationID    string `json:"conversationId,omitempty"`
	Type              string `json:"type,omitempty"`
	ConversationToken string `json:"conversationToken,omitempty"`
}

// Inputs ...
type Inputs struct {
	Intent    string       `json:"intent,omitempty"`
	RawInputs []*RawInputs `json:"rawInputs,omitempty"`
	Arguments []*Arguments `json:"arguments,omitempty"`
}

// RawInputs ...
type RawInputs struct {
	InputType string `json:"inputType,omitempty"`
	Query     string `json:"query,omitempty"`
}

// Arguments ...
type Arguments struct {
	Name      string `json:"name,omitempty"`
	RawText   string `json:"rawText,omitempty"`
	TextValue string `json:"textValue,omitempty"`
}

// Surface ...
type Surface struct {
	Capabilities []*Capabilities `json:"capabilities,omitempty"`
}

// Capabilities ...
type Capabilities struct {
	Name string `json:"name,omitempty"`
}

// AvailableSurfaces ...
type AvailableSurfaces struct {
	Capabilities []*Capabilities `json:"capabilities,omitempty"`
}

// ActionsResponse ...
type ActionsResponse struct {
	FulfillmentText    string              `json:"fulfillmentText,omitempty"`
	FollowupEventInput *FollowupEventInput `json:"followupEventInput,omitempty"`
	GPayload           *GPayload           `json:"payload,omitempty"`
}

// FollowupEventInput ...
type FollowupEventInput struct {
	Name string `json:"name,omitempty"`
}

// GPayload ...
type GPayload struct {
	Google *Google `json:"google,omitempty"`
}

// Google ...
type Google struct {
	ExpectUserResponse bool          `json:"expectUserResponse,omitempty"`
	RichResponse       *RichResponse `json:"richResponse,omitempty"`
	SystemIntent       *SystemIntent `json:"systemIntent,omitempty"`
}

// RichResponse ...
type RichResponse struct {
	Items []*Items `json:"items,omitempty"`
}

// Items ...
type Items struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
}

// SimpleResponse ...
type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
	Ssml         string `json:"ssml,omitempty"`
	DisplayText  string `json:"displayText,omitempty"`
}

// SystemIntent ...
type SystemIntent struct {
	Intent string `json:"intent,omitempty"`
}
