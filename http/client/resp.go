package client

import "github.com/imroc/req/v3"

func GetHeaderString(resp *req.Response) (headerString string) {
	headerMap := map[string]string{}
	for k := range resp.Header {
		if k != "Set-Cookie" {
			headerMap[k] = resp.Header.Get(k)
		}
	}
	for _, ck := range resp.Cookies() {
		headerMap["Set-Cookie"] += ck.String() + ";"
	}
	for k, v := range headerMap {
		headerString += k + ": " + v + "\n"
	}
	return headerString
}
