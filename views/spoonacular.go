package views

// GetCookingParams ...
type GetCookingParams struct {
	FoodName  string
	DietTypes string
}

// SpoonacularResult ...
type SpoonacularResult struct {
	Results      []*Results `json:"results"`
	Offset       int        `json:"offset,omitempty"`
	Number       int        `json:"number,omitempty"`
	TotalResults int        `json:"totalResults,omitempty"`
}

// Results ...
type Results struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	ImageType string `json:"imageType"`
}
