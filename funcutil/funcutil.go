package funcutil

import (
	"reflect"
	"runtime"
)

func NameOfFunction(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
