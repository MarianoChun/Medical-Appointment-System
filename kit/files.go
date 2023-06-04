package kit

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	var filesPath []string = make([]string, 0)

	for i := 0; i < len(files); i++ {
		if !files[i].IsDir() && !strings.EqualFold(files[i].Name(), "") {
			filesPath = append(filesPath, fmt.Sprintf("%s/%s", path, files[i].Name()))
		}
	}

	return filesPath, nil
}
