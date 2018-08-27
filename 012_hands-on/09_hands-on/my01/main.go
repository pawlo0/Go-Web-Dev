package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type quote struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

func main() {
	csvFile, _ := os.Open("table.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var stock []quote
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		stock = append(stock, quote{line[0], line[1], line[2], line[3], line[4], line[5], line[6]})
		fmt.Println(line[1])
	}
}
