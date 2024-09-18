# Steve's SDK (Go)

A great celebrity once said: It's my go. Here we are, doing go with Steve Reed,
providing Gophers the incredible capability to call [Steve's API](https://github.com/zhufucdev/api.zhufucdev)
with no library overhead.

## Getting started
Go with me,
```shell
go install github.com/zhufucdev/sdk_go@v0.0.1
```
Now check for update with `UpdateQuery` and stuff,
```go
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
```