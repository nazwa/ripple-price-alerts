package watcher

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// This is a default ripple api
type WatcherStruct struct {
	URL           string `json:"url"`
	Notificaitons []struct {
		Service  string            `json:"service"`
		Settings map[string]string `json:"settings"`
	} `json:"notifications"`
	Pairs       []*Item
	AlertsFound bool
}

func NewWatcher(url string) *WatcherStruct {
	watcher := &WatcherStruct{
		URL: url,
	}

	return watcher
}

func (w *WatcherStruct) AddPair(item *Item) {
	w.Pairs = append(w.Pairs, item)
}

// Queries ripple api for the latest prices
func (w *WatcherStruct) Update() {

	for _, item := range w.Pairs {
		err := w.GetRate(item)
		item.Error = err
	}

	return
}

func (w *WatcherStruct) GetRate(item *Item) error {

	json, err := item.GetJson()
	if err != nil {
		return err
	}

	response, err := http.Post(w.URL, "application/json", bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = item.ParseResponse(contents); err != nil {
		return err
	}
	if item.CheckRate(item.Rate) {
		w.AlertsFound = true
	}

	return nil

}
