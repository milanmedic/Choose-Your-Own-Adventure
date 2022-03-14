package filereader

import (
	"io/ioutil"
	"path/filepath"
)

type FileReader struct{}

func CreateFileReader() *FileReader {
	return &FileReader{}
}

func (fr *FileReader) ReadFile(filename string) ([]byte, error) {
	path := getFilePath(filename)

	contents, err := ioutil.ReadFile(filepath.Join(path, filename))

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func getFilePath(filename string) string {
	return filepath.Dir(filename)
}
