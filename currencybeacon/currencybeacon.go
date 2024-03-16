package currencybeacon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type CurrencyBeacon struct {
	apiKey string
}

func (cb *CurrencyBeacon) get(path string, params url.Values, result interface{}) error {
	params.Add("api_key", cb.apiKey)

	u := url.URL{
		Scheme:   "https",
		Host:     "api.currencybeacon.com",
		Path:     path,
		RawQuery: params.Encode()}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &result)

	return nil
}

func (cb *CurrencyBeacon) GetRates() (json.RawMessage, error) {
	var response ratesResponse

	err := cb.get("/v1/latest", url.Values{"base": {"USD"}}, &response)
	if err != nil {
		return json.RawMessage{}, err
	}

	return response.Response.Rates, nil
}

func NewCurrencyBeaconFromEnv() *CurrencyBeacon {
	return &CurrencyBeacon{
		apiKey: os.Getenv("CURRENCYBEACON_API_KEY")}
}
