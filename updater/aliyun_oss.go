package updater

import (
	"fmt"
	"github.com/imroc/req/v3"
	"time"
)

type AliyunOSS struct {
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
}

func NewAliyunOSS(accessKeyID, accessKeySecret string) Updater {
	return &AliyunOSS{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	}
}

func (o AliyunOSS) Fetch(url, filename string) (err error) {
	client := req.C()
	fmt.Println("start downloading...")
	callback := func(info req.DownloadInfo) {
		if info.Response.Response != nil {
			fmt.Printf("downloaded %.2f%%\n", float64(info.DownloadedSize)/float64(info.Response.ContentLength)*100.0)
		}
	}
	_, err = client.R().
		SetOutputFile(filename).
		SetDownloadCallbackWithInterval(callback, 1*time.Second).
		Get(url)
	fmt.Println("finished")
	return
}
