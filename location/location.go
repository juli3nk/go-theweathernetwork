package location

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/go-utils/json"
)

type Config struct {
	Locations Locations
}

func Search(location string) (*Config, error) {
	url := "https://www.theweathernetwork.com/ca/api/location/search"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")

	q := req.URL.Query()
	q.Add("searchText", location)

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	locations := new(Locations)

	if err := json.Decode(body, locations); err != nil {
		return nil, err
	}

	return &Config{Locations: *locations}, nil
}

func (c *Config) Find(t, city, region string) (Location, error) {
	t2 := strings.ToLower(t)
	city2 := strings.ToLower(city)
	region2 := strings.ToLower(region)

	for _, loc := range c.Locations {
		if strings.ToLower(loc.Type) == t2 && utils.StripCtlAndExtFromUnicode(strings.ToLower(loc.Name)) == city2 && strings.ToLower(loc.ProvCode) == region2 {
			return loc, nil
		}
	}

	return Location{}, fmt.Errorf("Location is not found")
}
