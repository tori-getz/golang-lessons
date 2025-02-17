package files

import (
	"fmt"
	"os"
)

type JsonDatabase struct {
	fileName string
}

func NewJsonDatabase(fileName string) *JsonDatabase {
	return &JsonDatabase{
		fileName: fileName,
	}
}

func (db *JsonDatabase) Read() ([]byte, error) {
	file, err := os.ReadFile(db.fileName)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (db *JsonDatabase) Write(content []byte) {
	file, err := os.Create(db.fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return
	}
}
