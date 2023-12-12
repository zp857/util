package convert

import "golang.org/x/text/encoding/simplifiedchinese"

type Charset string

const (
	UTF8 = Charset("UTF-8")
	GBK  = Charset("GBK")
)

func Byte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GBK:
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
