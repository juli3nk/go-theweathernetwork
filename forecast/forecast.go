package forecast

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juli3nk/go-utils/json"
)

type API struct {
	Location string
	Lang     string
	Domain   string
	URL      string
}

func get(url string, query map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")

	if len(query) > 0 {
		q := req.URL.Query()

		for k, v := range query {
			q.Add(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func New(location, lang string) (API, error) {
	domain := "www.theweathernetwork.com"
	if lang == "fr" {
		domain = "www.meteomedia.com"
	}

	urlBase := fmt.Sprintf("https://%s/api", domain)

	return API{
		Location: location,
		Lang:     lang,
		Domain:   domain,
		URL:      urlBase,
	}, nil
}

//https://www.theweathernetwork.com/ca/api/savedlocation/index/?placecodes=CAQC0363
func (api *API) Current() (*CurrentArr, error) {
	url := fmt.Sprintf("%s/savedlocation/index/", api.URL)

	data, err := get(url, map[string]string{"placecodes": api.Location})
	if err != nil {
		return nil, err
	}

	current := new(CurrentArr)

	if err := json.Decode(data, current); err != nil {
		return nil, err
	}

	return current, nil
}

func (api *API) Today() (*WeatherText, error) {
	url := fmt.Sprintf("%s/weathertext/%s", api.URL, api.Location)

	data, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	wt := new(WeatherText)

	if err := json.Decode(data, wt); err != nil {
		return nil, err
	}

	return wt, nil
}

func (api *API) Get() (*TheWeatherNetwork, error) {
	url := fmt.Sprintf("%s/data/%s/cm", api.URL, api.Location)

	data, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	twn := new(TheWeatherNetwork)

	if err := json.Decode(data, twn); err != nil {
		return nil, err
	}

	return twn, err

}

func (api *API) Hourly() (*TheWeatherNetwork, error) {
	url := fmt.Sprintf("%s/data/%s/hourly/cm", api.URL, api.Location)

	data, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	twn := new(TheWeatherNetwork)

	if err := json.Decode(data, twn); err != nil {
		return nil, err
	}

	return twn, err

}
