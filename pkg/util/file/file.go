package file

import (
	"io/ioutil"
	"os"
)

func Read(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	return byteValue, nil
}
