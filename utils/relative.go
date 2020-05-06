package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var ProjectDir string
var ImagesDir string

func init() {
	initRelative()
}

func initRelative() {
	_, fileName, _, _ := runtime.Caller(0)
	ProjectDir = filepath.ToSlash(filepath.Dir(filepath.Dir(fileName)))
	ImagesDir = filepath.Join(ProjectDir, "images")
}

func relative(path string) string {
	return strings.TrimPrefix(filepath.ToSlash(path), ProjectDir)
}

func NewError(e interface{}) error {
	if e != nil {
		_, fn, line, _ := runtime.Caller(1)
		return fmt.Errorf("[error] %s:%d %v", relative(fn), line, e)
	}
	return nil
}
