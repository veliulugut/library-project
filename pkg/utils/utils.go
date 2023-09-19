package utils

import (
	"context"
	"encoding/csv"
	"fmt"
	"library/ent"
	"log"
	"math/rand"
	"os"
)

func CreateTestData(db *ent.Client) error {
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

	for _, row := range records[1:] {
		_, err := db.Book.Create().SetTitle(row[0]).SetAuthor(row[1]).
			SetGenre(row[2]).SetHeight(row[3]).SetPublisher(row[4]).Save(context.Background())

		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("book data :%w", err)
		}

	}
	log.Println("test data create")
	return nil

}

func GenerateRandomOTPCode() string {
	chars := "ABCDEFGHJKMNPQRSTUVWXYZ123456789"

	result := ""

	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(chars))
		result += string(chars[randomIndex])
	}

	return result
}
