package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("data2.dat")
	if err != nil {
		log.Fatalln(err)
	}
	err2 := ioutil.WriteFile("data2-copy.txt", fileContents, 0x777)
	if err2 != nil {
		log.Fatalln(err2)
	}
	fmt.Println("file copied")
}
