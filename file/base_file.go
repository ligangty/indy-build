package file

import (
	"os"
	"strings"
)

func GetTempDir() string {
	temp := os.Getenv("TMPDIR")
	if strings.TrimSpace(temp) == "" {
		temp = "/tmp"
	}
	return temp
}

func StoreFile(fileName string, content string) {
	exists := FileExists(fileName)

	var f *os.File
	var err error
	if exists {
		os.Remove(fileName)
	}
	f, err = os.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
