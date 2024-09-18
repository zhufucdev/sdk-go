package sdk

type OperatingSystem int

const (
	NoOs OperatingSystem = iota
	Android
	MacOS
	Linux
	Windows
)

func (os OperatingSystem) String() string {
	return [...]string{"", "android", "darwin", "linux", "windows"}[os]
}

type Architecture int

const (
	NoArch Architecture = iota
	Universal
	Arm64
	Arm32
	X86
	X86_64
)

func (a Architecture) String() string {
	return [...]string{"", "universal", "arm64", "arm32", "x86", "x86_64"}[a]
}

type UpdateQuery struct {
	Product        string
	Os             OperatingSystem
	CurrentVersion string
	Arch           Architecture
}

type ReleaseAsset struct {
	VersionName string `json:"versionName,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Extension   string `json:"extension,omitempty"`
	Url         string `json:"url,omitempty"`
}

type ProductQuery struct {
	Name     string   `json:"name,omitempty"`
	Key      string   `json:"key,omitempty"`
	Category []string `json:"category,omitempty"`
}
