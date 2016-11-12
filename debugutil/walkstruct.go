package debugutil

import (
	"fmt"
	"reflect"
)

func walk(t interface{}) string {
	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()
	buf := ""
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		buf += fmt.Sprintf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	return buf
}
