package iputil

import (
	"fmt"
	"math/big"
	"net"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

var (
	SpecialCIDRs = sync.Map{}
)

func GetLookupIpFromUrL(rawUrl string) (ips []net.IP, err error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return ips, err
	}
	host := u.Hostname()
	ips, err = net.LookupIP(host)
	if err != nil {
		return ips, err
	}
	return ips, nil
}

// ExtractPortFromUrl 从 url 中提取 port 信息
func ExtractPortFromUrl(rawUrl string) (port string, err error) {
	isIpv6 := isIPv6(rawUrl)
	if isIpv6 {
		port, err = getPortFromIPv6(rawUrl)
	} else {
		port, err = getPortFromCommonUrl(rawUrl)
	}

	return port, nil
}

// 判断 URL 是否是 ipv6
func isIPv6(rawUrL string) bool {
	trimUrl := strings.TrimPrefix(rawUrL, "http://")
	trimUrl = strings.TrimPrefix(trimUrl, "https://")
	trimUrl = strings.TrimPrefix(trimUrl, "https:")
	trimUrl = strings.TrimPrefix(trimUrl, "http:")
	host, _, err := net.SplitHostPort(trimUrl)
	if err != nil {
		host = trimUrl
	}
	ip := net.ParseIP(host)
	return ip != nil && strings.Contains(host, ":")
}

func getPortFromIPv6(url string) (string, error) {
	// 使用正则表达式从 URL 中提取端口号
	re := regexp.MustCompile(`]:(\d+)$`)
	match := re.FindStringSubmatch(url)
	if len(match) < 2 {
		return "", fmt.Errorf("port not found in URL")
	}

	// 返回匹配的端口号
	return match[1], nil
}

func getPortFromCommonUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	port := ""
	if parsedUrl.Port() == "" {
		if parsedUrl.Scheme == "https" {
			port = "443"
		} else {
			port = "80"
		}
	} else {
		port = parsedUrl.Port()
	}
	return port, nil
}

func ValidInnerIp(ip string) bool {
	validInnerIps := []string{
		// 10.0.0.0/8
		"10.0.0.0-10.255.255.255",
		// 172.16.0.0/12
		"172.16.0.0-172.31.255.255",
		// 192.168.0.0/16
		"192.168.0.0-192.168.255.255",
	}
	for _, v := range validInnerIps {
		ipSlice := strings.Split(v, `-`)
		if len(ipSlice) < 0 {
			continue
		}
		if InetAtoi(ip) >= InetAtoi(ipSlice[0]) && InetAtoi(ip) <= InetAtoi(ipSlice[1]) {
			return true
		}
	}
	var valid bool
	SpecialCIDRs.Range(func(key, value any) bool {
		t := net.ParseIP(ip)
		if t != nil {
			_, a, _ := net.ParseCIDR(key.(string))
			if a.Contains(t) {
				valid = true
				return false
			}
		}
		return true
	})
	return valid
}

func InetAtoi(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(strings.TrimSpace(ip)).To4())
	return ret.Int64()
}
