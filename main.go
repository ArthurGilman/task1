package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Product string
	Price   int
	Rating  int
}

func main() {
	var (
		p Data
		r Data
	)
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("No files")
	}
	for _, f := range args {
		var err error

		if strings.Contains(f, "csv") {
			p, r, err = GetDataFromCSV(f)
		} else if strings.Contains(f, "json") {
			p, r, err = GetDataFromJson(f)
		} else {
			log.Println("Inccorect file")
			continue
		}
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Most expensive : %s\nMost rated: %s\n", p.Product, r.Product)
	}

}

func GetDataFromCSV(filename string) (Data, Data, error) {
	var (
		maxP Data
		maxR Data
	)
	data := make([]Data, 0, 50)

	file, err := os.Open(filename)
	if err != nil {
		return Data{}, Data{}, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = 3

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		p, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		r, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}
		d := Data{Product: record[0], Price: p, Rating: r}
		data = append(data, d)
	}
	if len(data) != 0 {
		maxP = data[0]
		maxR = data[0]
	} else {
		return Data{}, Data{}, errors.New("There is no data in the file")
	}
	for _, d := range data {
		if d.Price > maxP.Price {
			maxP = d
		}
		if d.Rating > maxR.Rating {
			maxR = d
		}
	}
	return maxP, maxR, nil
}

func GetDataFromJson(filename string) (Data, Data, error) {
	var (
		maxP Data
		maxR Data
	)
	data := make([]Data, 0, 50)
	file, err := os.Open(filename)
	if err != nil {
		return Data{}, Data{}, err
	}
	defer file.Close()

	for {
		d := []Data{}
		reader := bufio.NewReader(file)
		decoder := json.NewDecoder(reader)
		if err := decoder.Decode(&d); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for _, val := range d {
			data = append(data, val)
		}
	}
	if len(data) != 0 {
		maxP = data[0]
		maxR = data[0]
	} else {
		return Data{}, Data{}, errors.New("There is no data in the file")
	}
	for _, d := range data {
		if d.Price > maxP.Price {
			maxP = d
		}
		if d.Rating > maxR.Rating {
			maxR = d
		}
	}

	return maxP, maxR, nil
}
