package watcher

import (
	"encoding/json"
	"fmt"
)

type AlertDirection string

const (
	AlertAbove AlertDirection = "above"
	AlertBelow AlertDirection = "below"
)

type CurrencyStruct struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer,omitempty"`
}

type Item struct {
	Base    CurrencyStruct `json:"base"`
	Counter CurrencyStruct `json:"counter"`
	Live    bool           `json:"live"`
	Alert   bool           `json:"-"`

	Rate          float64        `json:"rate,omitempty"`
	AlertRate     float64        `json:"-"`
	AlertMode     AlertDirection `json:"-"`
	AlertReported bool           `json:"-"`

	Error error `json:"-"`

	Value float64        `json:"value,omitempty"`
	Mode  AlertDirection `json:"mode,omitempty"`
}

// Parses json string and populates a new item object
func ItemFromJson(data []byte) (*Item, error) {
	i := &Item{}

	if err := json.Unmarshal(data, i); err != nil {
		return nil, err
	}

	switch i.Mode {
	case AlertAbove, AlertBelow:
		i.AlertMode = i.Mode
	default:
		return nil, fmt.Errorf("Config load failed: %s is not a valid mode.", i.Mode)
	}
	i.AlertRate = i.Value

	i.Value = 0
	i.Mode = ""

	return i, nil

}

func (i *Item) GetJson() ([]byte, error) {
	return json.Marshal(i)
}
func (i *Item) ParseResponse(response []byte) error {
	i.Rate = 0

	var temp []Item

	err := json.Unmarshal(response, &temp)
	if err != nil {
		return err
	}

	if len(temp) == 0 {
		return fmt.Errorf("Empty response: %x", response)
	}
	i.Rate = temp[0].Rate

	return nil
}

func (i *Item) CheckRate(rate float64) bool {
	i.Alert = false
	if i.AlertMode == AlertAbove && rate > i.AlertRate {
		i.Alert = true
	}
	if i.AlertMode == AlertBelow && rate < i.AlertRate {
		i.Alert = true
	}
	if i.AlertReported && !i.Alert {
		i.AlertReported = false
	}

	return i.Alert
}

func (i *Item) NewAlertDetected() bool {
	return !i.AlertReported && i.Alert
}

func (i *Item) MarkReported() {
	i.AlertReported = true
}

func (i *Item) String() string {

	return fmt.Sprintf("%s->%s: %f; Marker set: \"%s\" %f", i.Base.Currency, i.Counter.Currency, i.Rate, i.AlertMode, i.AlertRate)
}
