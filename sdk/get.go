package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	u "net/url"
)

func GetReleaseAsset(baseUrl string, query UpdateQuery) (ReleaseAsset, error) {
	url, err := u.Parse(fmt.Sprintf("%s/release", baseUrl))
	if err != nil {
		return ReleaseAsset{}, err
	}
	queryValues := make(u.Values)
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

	httpResponse, err := http.Get(url.String())
	if err != nil {
		return ReleaseAsset{}, err
	}

	decoder := json.NewDecoder(httpResponse.Body)
	var result ReleaseAsset
	err = decoder.Decode(&result)
	if err != nil {
		return ReleaseAsset{}, err
	}
	return result, nil
}
