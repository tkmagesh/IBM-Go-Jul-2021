/* using the bufio.NewReader to read sentences from the source */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("data1.dat")
	if fileError != nil {
		log.Fatalln(fileError)
	}
	defer fileHandle.Close()
	sentenceCount := 0
	inputReader := bufio.NewReader(fileHandle)
	for {
		sentence, err := inputReader.ReadString('.')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		sentenceCount++
		fmt.Println("Sentence : ", sentenceCount)
		fmt.Println(sentence)
	}
}
