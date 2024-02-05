package main

import (
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
	"parser/internal/db"
	"parser/internal/utils"
)

var Version=""

func main() {
	init()

	files, err := os.ReadDir("./data/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileName := f.Name()
		file, err := os.Open("./data/" + fileName)
		if err != nil {
			log.Fatal(err)
		}

		fileContent, _ := io.ReadAll(file)
		content := utils.ReadLines(fileContent)

		for _, line := range content {
			conn := db.NewConnection(err)

			_, err = conn.Exec(
				fmt.Sprintf(
					"INSERT INTO devices (title, ip, totalBytes) VALUES (%s, %s, %d)",
					line.Title,
					line.Ip,
					utils.calc(line.InputBytes, line.OutputBytes),
				),
			)

			if err != nil {
				log.Fatalf("Database insert failed: %v", err)
			}
		}
		fmt.Printf("Insert file %s content into database successful\n", fileName)
	}
}

func init() {
	if (Version == ""){
		log.Fatalf("Version is not provided")
	}
}
