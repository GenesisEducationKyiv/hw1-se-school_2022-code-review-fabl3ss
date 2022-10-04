package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAndParseBody(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		err = r.Body.Close()
	}()

	err = json.NewDecoder(r.Body).Decode(target)
	return err
}

func DoHttpAndParseBody(req *http.Request, target interface{}) error {
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err = r.Body.Close()
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err

	}

	return json.Unmarshal(body, target)
}
