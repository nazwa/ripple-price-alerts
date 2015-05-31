package watcher

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigStruct struct {
	Watcher *WatcherStruct    `json:"watcher"`
	Jobs    []json.RawMessage `json:"jobs"`
}

func LoadFromFile(path string) (*WatcherStruct, error) {
	jobs_file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &ConfigStruct{}

	err = json.Unmarshal(jobs_file, &cfg)
	if err != nil {
		return nil, err
	}

	for _, data := range cfg.Jobs {
		item, err := ItemFromJson(data)
		if err != nil {
			return nil, err
		}
		cfg.Watcher.Pairs = append(cfg.Watcher.Pairs, item)
	}

	return cfg.Watcher, nil
}
