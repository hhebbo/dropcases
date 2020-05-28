package fileCache

import (
	"io/ioutil"
	"os"
)

const (
	FILECACHE_ROOT_PATH    = "storage/fileCache/"
	FILECACHE_WRITING_MODE = 0644
)

func Get(filePath string) string {
	value, err := ioutil.ReadFile(FILECACHE_ROOT_PATH + filePath)
	if err != nil {
		panic(err)
	}

	return string(value)
}

func Save(filePath string, value string) {
	err := ioutil.WriteFile(FILECACHE_ROOT_PATH+filePath, []byte(value), FILECACHE_WRITING_MODE)
	if err != nil {
		panic(err)
	}
}

func DirectoryExistsOrCreate(path string) bool {
	if _, err := os.Stat(FILECACHE_ROOT_PATH + path); os.IsNotExist(err) {
		os.MkdirAll(FILECACHE_ROOT_PATH+path, os.ModePerm)

		return false
	}

	return true
}
