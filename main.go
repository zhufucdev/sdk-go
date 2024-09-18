package main

import (
	"fmt"
	"github.com/zhufucdev/sdk-go/sdk"
)

func main() {
	q := sdk.UpdateQuery{
		Product:        "turret",
		Os:             sdk.Android,
		CurrentVersion: "1.0.0",
	}
	a, err := sdk.GetReleaseAsset("http://localhost:3000/api", q)
	if err != nil {
		fmt.Println(err.Error())
	} else if a != nil {
		fmt.Println(a.VersionName, a.Url)
	} else {
		fmt.Println("No update found.")
	}
}
