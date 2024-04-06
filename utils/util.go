package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
)

func PasswordEncrypt(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func Struct2Map[T any](str T) map[string]any {
	v := reflect.ValueOf(str)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	m := make(map[string]any)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		value := reflect.Zero(f.Type())
		if !f.CanInterface() || value.Interface() == f.Interface() {
			continue
		}
		m[t.Field(i).Name] = f.Interface()
	}
	return m
}
