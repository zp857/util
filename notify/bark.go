package notify

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type Bark struct {
	key       string
	reqClient *req.Client
}

func NewBark(key string) Notifyer {
	return &Bark{
		key:       key,
		reqClient: req.C(),
	}
}

func (b Bark) Send(msg Msg) (body string, err error) {
	url := fmt.Sprintf("https://api.day.app/%v", b.key)
	r := b.reqClient.R()
	r.SetBody(msg)
	var resp *req.Response
	resp, err = r.Post(url)
	if err != nil {
		return
	}
	body = resp.String()
	return
}
