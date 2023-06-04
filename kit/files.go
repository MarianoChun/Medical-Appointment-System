package kit

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetFilesPathByFolder(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var filesPath []string = make([]string, len(files))

	for i := 0; i < len(files); i++ {
		if !files[i].IsDir() {
			filesPath = append(filesPath, files[i].Name())
		}
	}

	return filesPath, nil
}
