package updater

type Updater interface {
	Fetch(url, filename string) (err error)
}

type Config struct {
	AliyunOSS AliyunOSS `json:"aliyunOSS"`
}
