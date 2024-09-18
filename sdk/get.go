package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	u "net/url"
)

func GetReleaseAsset(baseUrl string, query UpdateQuery) (*ReleaseAsset, error) {
	url, err := u.Parse(fmt.Sprintf("%s/release", baseUrl))
	if err != nil {
		return nil, err
	}
	queryValues := make(u.Values)
	if query.Product == "" {
		return nil, errors.New("product required")
	}
	queryValues.Set("product", query.Product)
	if query.Arch != NoArch {
		queryValues.Set("arch", query.Arch.String())
	}
	if query.Os != NoOs {
		queryValues.Set("os", query.Os.String())
	}
	if query.CurrentVersion != "" {
		queryValues.Set("current", query.CurrentVersion)
	}
	url.RawQuery = queryValues.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 || res.StatusCode < 200 {
		if res.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("http error: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	var result ReleaseAsset
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
