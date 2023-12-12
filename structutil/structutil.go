package structutil

import (
	"bytes"
	"encoding/json"
	"strings"
)

func JsonMarshalIndent(obj interface{}) (out string) {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false) // 不转义特殊字符
	encoder.SetIndent("", "  ")
	err := encoder.Encode(obj)
	if err != nil {
		return ""
	}
	out = byteBuf.String()
	out = strings.TrimRight(out, "\n")
	return
}
