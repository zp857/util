package client

import (
	"crypto/tls"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

type Options struct {
	Proxy   string   `yaml:"proxy" json:"proxy"`
	Timeout int      `yaml:"timeout" json:"timeout"`
	Headers []string `yaml:"headers" json:"headers"`
}

func NewReqClient(options *Options) *req.Client {
	reqClient := req.C().EnableDumpEachRequest()
	reqClient.GetTLSClientConfig().InsecureSkipVerify = true
	reqClient.SetCommonHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
	})
	reqClient.GetTLSClientConfig().MinVersion = tls.VersionTLS10
	reqClient.SetRedirectPolicy(req.AlwaysCopyHeaderRedirectPolicy("Cookie"))
	if options.Proxy != "" {
		reqClient.SetProxyURL(options.Proxy)
	}
	var key, value string
	for _, header := range options.Headers {
		tokens := strings.SplitN(header, ":", 2)
		if len(tokens) < 2 {
			continue
		}
		key = strings.TrimSpace(tokens[0])
		value = strings.TrimSpace(tokens[1])
		reqClient.SetCommonHeader(key, value)
	}
	if options.Timeout != 0 {
		reqClient.SetTimeout(time.Duration(options.Timeout) * time.Second)
	} else {
		reqClient.SetTimeout(10 * time.Second)
	}
	return reqClient
}
