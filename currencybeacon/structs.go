package currencybeacon

import (
	"encoding/json"
)

type ratesResponse struct {
	Response struct {
		Date  string          `json:"date"`
		Rates json.RawMessage `json:"rates"`
	} `json:"response"`
}
