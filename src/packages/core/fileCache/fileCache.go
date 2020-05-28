package fileCache

import (
	"io/ioutil"
	"os"
)

const (
	FILECACHE_PATH = "storage/fileCache/"
)

func Get(filePath string) string {
	value, err := ioutil.ReadFile(FILECACHE_PATH + filePath)
	if err != nil {
		panic(err)
	}

	return string(value)
}

func Save(filePath string, value string) {
	err := ioutil.WriteFile(FILECACHE_PATH+filePath, []byte(value), 0644)
	if err != nil {
		panic(err)
	}
}

func Exists(filePath string) bool {
	if _, err := os.Stat(FILECACHE_PATH + filePath); err == nil {
		value := Get(filePath)
		if value != "" {
			return true
		}
	}

	return false
}
