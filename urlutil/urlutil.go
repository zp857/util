package urlutil

import (
	"github.com/zp857/util/sliceutil"
	"net/url"
	"path/filepath"
	"strings"
)

func CountUniURLFile(links []string) int {
	var urls []string
	for _, link := range links {
		u, err := url.Parse(link)
		if err != nil {
			continue
		}
		u.RawQuery = ""
		urls = append(urls, u.String())
	}
	urls = sliceutil.Unique(urls)
	return len(urls)
}

func GetFileName(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	file := filepath.Base(u.Path)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	return file
}
