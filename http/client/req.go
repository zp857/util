package client

import (
	"crypto/tls"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

type Options struct {
	DumpAll bool     `yaml:"dumpAll" json:"dumpAll"`
	Proxy   string   `yaml:"proxy" json:"proxy"`
	Timeout int      `yaml:"timeout" json:"timeout"`
	Headers []string `yaml:"headers" json:"headers"`
}

const (
	defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	cookieHeader     = "cookie"
)

func NewReqClient(options *Options) *req.Client {
	reqClient := req.C().EnableDumpEachRequest()
	if options.DumpAll {
		reqClient.EnableDumpAll()
	}
	reqClient.GetTLSClientConfig().InsecureSkipVerify = true
	reqClient.SetCommonHeaders(map[string]string{
		"User-Agent": defaultUserAgent,
	})
	reqClient.GetTLSClientConfig().MinVersion = tls.VersionTLS10
	reqClient.SetRedirectPolicy(req.AlwaysCopyHeaderRedirectPolicy(cookieHeader))
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

var (
	toHttps = []string{
		"sent to HTTPS port",
		"This combination of host and port requires TLS",
		"Instead use the HTTPS scheme to",
		"This web server is running in SSL mode",
	}
)

func FirstGet(client *req.Client, url string) (resp *req.Response, err error) {
	request := client.R()
	var scheme string
	var flag bool
	if !strings.HasPrefix(url, "http") {
		scheme = "http://"
		resp, err = request.Get(scheme + url)
		if err != nil {
			scheme = "https://"
			flag = true
		} else {
			for _, str := range toHttps {
				if strings.Contains(resp.String(), str) {
					scheme = "https://"
					flag = true
					break
				}
			}
		}
	} else if strings.HasPrefix(url, "http://") {
		resp, err = request.Get(url)
		if err != nil {
			scheme = "https://"
			url = url[7:]
			flag = true
		} else {
			for _, str := range toHttps {
				if strings.Contains(resp.String(), str) {
					scheme = "https://"
					url = url[7:]
					flag = true
				}
			}
		}
	} else {
		flag = true
	}
	if flag {
		resp, err = request.Get(scheme + url)
	}
	return
}
