package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeHTTPRequest(url string) error {
	var urls []string
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(body, &urls)

	fmt.Println("Response Body:", urls)

	return nil

}
