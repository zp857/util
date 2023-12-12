package crypto

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
)

func Encrypt(rawData, key []byte) (string, error) {
	dst, err := openssl.AesECBEncrypt(rawData, key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}

func Decrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dst, err := openssl.AesECBDecrypt(data, key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}
