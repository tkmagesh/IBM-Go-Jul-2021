/* using the bufio.Scanner to read lines from the source */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("data2.dat")
	if fileError != nil {
		log.Fatalln(fileError)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	lineCount := 0
	for scanner.Scan() {
		str := scanner.Text()
		lineCount++
		fmt.Println("Line : ", lineCount)
		fmt.Println(str)
	}
}
