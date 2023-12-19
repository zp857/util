package stringutil

import "strings"

// SplitString 函数用于将字符串按照指定长度分割为切片
func SplitString(s string, n int) []string {
	var chunks []string
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func SplitItems(itemString string) []string {
	items := strings.Split(itemString, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}
