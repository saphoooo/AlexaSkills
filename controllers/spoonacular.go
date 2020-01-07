package controllers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"cooking.io/views"

	"github.com/pkg/errors"
)

// GetCookingInstructionIntent ...
func GetCookingInstructionIntent(p *views.GetCookingParams) ([]byte, error) {
	baseURL, err := url.Parse("https://api.spoonacular.com/recipes/complexSearch?")
	if err != nil {
		return nil, errors.WithMessage(err, "malformed URL")
	}
	params := url.Values{}
	params.Add("apiKey", os.Getenv("SPOONACULAR_APIKEY"))
	params.Add("number", "3")
	params.Add("offset", "0")
	params.Add("instructionsRequired", "true")
	if p.FoodName != "" {
		params.Add("query", p.FoodName)
	}
	if p.DietTypes != "" {
		params.Add("diet", p.DietTypes)
	}
	baseURL.RawQuery = params.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	js, err := ioutil.ReadAll(res.Body)
	return js, nil
}

// NewGetCookingParams ...
func NewGetCookingParams() *views.GetCookingParams {
	return &views.GetCookingParams{
		FoodName:  "",
		DietTypes: "",
	}
}
