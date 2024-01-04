package fileutil

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)

// PathExists 判断路径是否存在
func PathExists(path string) bool {
	f, flag := IsExists(path)
	if !flag {
		return false
	}
	return f.IsDir()
}

func FileExists(path string) bool {
	f, flag := IsExists(path)
	if !flag {
		return false
	}
	return !f.IsDir()
}

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

// WriteJSON 结果写入 json, 用于自测
func WriteJSON(filename string, data interface{}) (err error) {
	var bytes []byte
	bytes, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	err = WriteFile(filename, string(bytes))
	return
}

// WriteFile 写入文件
func WriteFile(filename string, data string) (err error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return
	}
	err = writer.Flush()
	return
}

func WritePath(path string, data string) (err error) {
	b := FileExists(path)
	var f *os.File
	if b {
		f, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	} else {
		_ = os.MkdirAll(filepath.Dir(path), 0666)
		f, err = os.Create(path)
	}
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.WriteString(data)
	return
}

// GetAllFile 获取指定目录下所有文件
func GetAllFile(dirPath string) (results []string, err error) {
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			results = append(results, path)
		}
		return nil
	})
	return
}

// ReadLines 按行读取文件
func ReadLines(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}
