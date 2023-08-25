package main

import (
	"bufio"
	"log"
	"os"
)

func ScanForSpecialized(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[0:i], nil
		}
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}
func main() {
	file, err := os.Open("buf.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanForSpecialized)

	for scanner.Scan() {
		println(scanner.Text())
	}

}
