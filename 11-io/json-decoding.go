package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First string `json:"first_name"`
	Last  string `json:"last_name"`
}

type Book struct {
	Title       string    `json:"book_title"`
	PageCount   int       `json:"pages"`
	ISBN        string    `json:"isbn"`
	Authors     []Name    `json:"authors"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"pub_date"`
}

func main() {
	file, err := os.Open("book.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var book Book
	dec := json.NewDecoder(file)
	if err := dec.Decode(&book); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(book)
}
