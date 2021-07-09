package main

import (
	"fmt"
	"io"
	"os"
)

type AlphaReader struct {
	Src io.Reader
}

func (alphaReader AlphaReader) Read(buffer []byte) (n int, err error) {

	inputData := make([]byte, len(buffer))
	count, err := alphaReader.Src.Read(inputData)
	if err != nil {
		return count, err
	}
	dataCount := 0
	for i := 0; i < len(inputData); i++ {
		if (inputData[i] >= 'A' && inputData[i] <= 'Z') || (inputData[i] >= 'a' && inputData[i] <= 'z') {
			buffer[dataCount] = inputData[i]
			dataCount += 1
		}
	}
	return dataCount, io.EOF
}

/* func copy(writer io.Writer, reader io.Reader) (int, error) {
	totalBytesRead := 0
	for {
		buffer := make([]byte, 1024)
		bytesRead, err := reader.Read(buffer)
		totalBytesRead += bytesRead
		if bytesRead > 0 {
			_, writeErr := writer.Write(buffer)
			if writeErr != nil {
				return totalBytesRead, writeErr
			}
		}
		if err != nil {
			return totalBytesRead, err
		}
	}
} */

func main() {
	//stringReader := strings.NewReader("Hi, How are you?")
	fileReader, _ := os.Open("data1.dat")
	alphaReader := AlphaReader{Src: fileReader}
	io.Copy(os.Stdout, alphaReader)
	fmt.Println("Done")
}
