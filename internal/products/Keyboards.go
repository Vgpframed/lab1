package products

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)
type Keyboard Product

func (k *Keyboard) GetProduct(filename string) *[]Keyboard {

	var inc = 0
	cnt := k.GetCount(filename)
	Products := make([]Keyboard, *cnt)
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
		if record[0] == "1"{
			k.Flag = record[0]
			k.Name = record[1]
			k.Type = record[2]
			k.Rate, _ = strconv.ParseInt(record[3], 0, 64)
			k.Price, _ = strconv.ParseInt(record[4], 0, 64)
			Products[inc] = *k
			inc++
		}
	}
	//fmt.Println(Products)
	return &Products
}

func (k *Keyboard) GetCount(filename string) *int {
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
		if record[0] == "1"{
			inc++
		}
	}
	return &inc
}
