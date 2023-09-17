package utils

import (
	"context"
	"encoding/csv"
	"fmt"
	"library/ent"
	"log"
	"os"
)

type Book struct {
	db *ent.Client
}

func NewBookData(db *ent.Client) *Book {
	return &Book{
		db: db,
	}
}

// Title,Author,Genre,Height,Publisher
type BookModel struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Height    string `json:"height"`
	Publisher string `json:"publisher"`
}

func (a *Book) BookData() error {
	fd, err := os.Open("books.csv")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("book data / open :%w", err)
	}

	log.Println("Successfully opened the CSV file")
	defer fd.Close()

	fileReader := csv.NewReader(fd)

	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("book data / read all :%w", err)
	}

	bookList := make([]BookModel, 0, len(records)-1)
	for _, row := range records[1:] {

		bookList = append(bookList, BookModel{
			Title:     row[0],
			Author:    row[1],
			Genre:     row[2],
			Height:    row[3],
			Publisher: row[4],
		})

	}

	for i := range bookList {
		result, err := a.db.Book.Create().SetTitle(bookList[i].Title).SetAuthor(bookList[i].Author).
			SetGenre(bookList[i].Genre).SetHeight(bookList[i].Height).SetPublisher(bookList[i].Publisher).Save(context.Background())

		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("book data :%w", err)
		}

		fmt.Println(result)
	}

	return nil

}
