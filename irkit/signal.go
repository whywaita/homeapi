package irkit

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/whywaita/yayoi/vars"
)

const (
	IrkitURL = "https://api.getirkit.com/1/messages"
)

func SendSignal(message string) error {
	values := url.Values{}
	values.Set("clientkey", vars.ClientKey)
	values.Add("deviceid", vars.DeviceID)
	values.Add("message", message)

	req, err := http.NewRequest(
		"POST",
		IrkitURL,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
