package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"cooking.io/spoonacular"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

// getCookingInstructionIntent makes a complex search on spoonacular API
// see https://spoonacular.com/food-api/docs#Search-Recipes-Complex
func getCookingInstructionIntent(p *spoonacular.GetCookingParams) ([]byte, error) {
	baseURL, err := url.Parse("https://api.spoonacular.com/recipes/complexSearch?")
	if err != nil {
		return nil, errors.WithMessage(err, "malformed URL")
	}
	params := url.Values{}
	params.Add("apiKey", os.Getenv("SPOONACULAR_APIKEY"))
	params.Add("number", "3")
	params.Add("offset", p.Offset)
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
	resp, err := ioutil.ReadAll(res.Body)
	return resp, nil
}

// newGetCookingParams is an helper function to initialize a GetCookingParams struct
func newGetCookingParams() *spoonacular.GetCookingParams {
	return &spoonacular.GetCookingParams{
		FoodName:  "",
		DietTypes: "",
		Offset:    "0",
	}
}

// resultsToText formats the results to create a single string
func resultsToText(results []byte) (string, error) {
	returnedString := "I found following dishes that you can cook"
	var r spoonacular.Result
	err := json.Unmarshal(results, &r)
	if err != nil {
		return "", errors.WithMessage(err, "unable to unmarshal spoonacular results")
	}
	for _, value := range r.Results {
		returnedString = returnedString + ", " + value.Title
	}
	return returnedString, nil
}

// keepCookingParams stores query parameters in redis for 120 seconds
func keepCookingParams(pool *redis.Pool, key string, params *spoonacular.GetCookingParams) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HMSET", key, "FoodName", params.FoodName, "DietTypes", params.DietTypes, "Offset", params.Offset)
	if err != nil {
		return errors.WithMessage(err, "error while inserting params in Redis")
	}
	_, err = conn.Do("EXPIRE", key, 120)
	if err != nil {
		return errors.WithMessage(err, "error while setting key expiry in Redis")
	}
	return nil
}

// getCookingParams retrieves query paramaters in redis
func getCookingParams(pool *redis.Pool, key string) (*spoonacular.GetCookingParams, error) {
	conn := pool.Get()
	defer conn.Close()

	exists, err := redis.Int(conn.Do("EXISTS", key))
	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, errors.New("ErrNoKey")
	}

	reply, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return &spoonacular.GetCookingParams{
		FoodName:  reply["FoodName"],
		DietTypes: reply["DietTypes"],
		Offset:    reply["Offset"],
	}, nil
}
