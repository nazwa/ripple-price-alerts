package watcher

import (
	"fmt"
	"github.com/sfreiberg/gotwilio"
	"strings"
)

func (w *WatcherStruct) Notify(item *Item) error {

	for _, service := range w.Notificaitons {

		switch strings.ToLower(service.Service) {
		case "twilio":
			twilio := gotwilio.NewTwilioClient(service.Settings["sid"], service.Settings["token"])

			_, exception, err := twilio.SendSMS(service.Settings["from"], service.Settings["to"], item.String(), "", "")

			if exception != nil {
				return fmt.Errorf("%v", exception)
			}
			if err != nil {
				return err
			}
		}
	}

	return nil

}
