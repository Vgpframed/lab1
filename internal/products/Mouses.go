package products

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)
type Mouse Product

func (m *Mouse) GetProduct(filename string) *[]Mouse {

	var inc = 0
	cnt := m.GetCount(filename)
	Products := make([]Mouse, *cnt)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		record, e := reader.Read()
		if e == io.EOF {
			break
		}
		if record[0] == "0"{
			m.Flag = record[0]
			m.Name = record[1]
			m.Type = record[2]
			m.Rate, _ = strconv.ParseInt(record[3], 0, 64)
			m.Price, _ = strconv.ParseInt(record[4], 0, 64)
			Products[inc] = *m
			inc++
		}
	}
	fmt.Println(Products)
	return &Products
}

func (m *Mouse) GetCount(filename string) *int {
	inc := 0
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	for {
		record, e := reader.Read()
		if e == io.EOF {
			break
		}
		if record[0] == "0"{
			inc++
		}
	}
	return &inc
}
