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
	book := Book{
		Title:     "My Book",
		PageCount: 375,
		ISBN:      "9781784395438",
		Authors:   []Name{{"Magesh", "Kuppan"}, {"Suresh", "Kannan"}, {"Ramesh", "Jayaraman"}},
		PublishDate: time.Date(
			2016, time.July,
			0, 0, 0, 0, 0, time.UTC),
	}

	file, err := os.Create("book.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := json.NewEncoder(file)
	if err := enc.Encode(book); err != nil {
		fmt.Println(err)
	}
}
