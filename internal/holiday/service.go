package holiday

import (
	"coding/session/internal/holiday/internal/api"
	"context"
	"fmt"
)

type holiday struct {
	Day     string
	WeekDay string
	Month   string
	Name    string
}

// holidaysFor retrieves holiday information for a specified country.
func holidaysFor(ctx context.Context, countryCode string) ([]holiday, error) {
	// Tasks:
	// 1. The first step will be removing the fmt.Println :D
	// 2. Let's use iterators to loop over the dates and add them to the holidays slice.
	// 3. We need to save the holiday information in a map that will hold the results of previous holiday requests.

	dates, err := api.Holidays(ctx, countryCode)
	if err != nil {
		return nil, err
	}

	var holidays []holiday
	fmt.Println(dates)

	return holidays, nil
}
