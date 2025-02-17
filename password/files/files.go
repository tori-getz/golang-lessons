package files

import (
	"fmt"
	"os"
)

func Read() {

}

func Write(fileName string, content []byte) {
	file, err := os.Create(fileName)

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
