package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type date struct {
	StringDate  string `json:"date"`
	LocalName   string `json:"localName"`
	CountryCode string `json:"countryCode"`
}

// Holidays fetches public holidays for a given country using the Nager.Date API.
// It takes a context for request cancellation and a country code as input,
// returning a slice of dates and an error in case there is any.
func Holidays(ctx context.Context, countryCode string) ([]date, error) {
	baseURL := fmt.Sprintf("https://date.nager.at/api/v3/publicholidays/%d/%s",
		time.Now().Year(),
		countryCode,
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var dates []date
	if err := json.NewDecoder(res.Body).Decode(&dates); err != nil {
		return nil, err
	}

	return dates, nil
}
