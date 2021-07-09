package main

import (
	"fmt"
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(buffer []byte) (n int, err error) {
	fmt.Println("size of the buffer = ", len(buffer))
	count := 0
	for i := 0; i < len(alphaReader); i++ {
		if (alphaReader[i] >= 'A' && alphaReader[i] <= 'Z') || (alphaReader[i] >= 'a' && alphaReader[i] <= 'z') {
			buffer[count] = alphaReader[i]
			count += 1
		}
	}
	return count, io.EOF
}

func copy(writer io.Writer, reader io.Reader) (int, error) {
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
}

func main() {
	alphaReader := AlphaReader("Hi, How are you?")
	//copy(os.Stdout, alphaReader)
	io.Copy(os.Stdout, alphaReader)
	fmt.Println("Done")
}
