package headless

import (
	"fmt"
)

func GetHeaderString(headers map[string]interface{}) (headerString string) {
	for k, v := range headers {
		headerString += fmt.Sprintf("%v: %v\n", k, v)
	}
	return headerString
}
