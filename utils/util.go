package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"os"
	"path/filepath"
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

func IsExist(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
}

func MakeDir(p string) error {
	exist, err := IsExist(p)
	if err != nil {
		return err
	}
	if !exist {
		err := os.Mkdir(p, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ListDir(p string) []os.FileInfo {
	files := make([]os.FileInfo, 0)
	_ = filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info)
		}
		return err
	})
	return files
}
