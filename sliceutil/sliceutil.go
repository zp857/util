package sliceutil

import (
	"fmt"
	"github.com/zp857/util/structutil"
	"strings"
)

func Unique[T comparable](slice []T) []T {
	var result []T

	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range result {
			if v == result[j] {
				skip = false
				break
			}
		}
		if skip {
			result = append(result, v)
		}
	}

	return result
}

func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}

	return false
}

func SplitItems(itemString string) []string {
	items := strings.Split(itemString, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}

func Remove[T comparable](slice []T, target T) []T {
	var result []T
	for i := 0; i < len(slice); i++ {
		if slice[i] != target {
			result = append(result, slice[i])
		}
	}
	return result
}

func DebugPrint[T comparable](slice []T) {
	fmt.Printf("len => %v\n%v\n", len(slice), structutil.JsonMarshalIndent(slice))
}
