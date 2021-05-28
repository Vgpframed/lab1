package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"

	"fmt"
	product "lab1/internal/products"
	"os"

)

var Log *os.File
func init(){
	var err error
	Log, err = os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateFileOpenFileWrite(filename string, network *bytes.Buffer) {

	_, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0777)
	if err != nil{
		fmt.Println("Не удалось открыть файл")
		return
	}
	ioutil.WriteFile(filename, network.Bytes(), os.ModePerm)

}
func enc (filename string){

	var prod product.Keyboard
	products := prod.GetProduct(Filename)
	network:= bytes.Buffer{}
	enc := gob.NewEncoder(&network)
	err := enc.Encode(products)
	if err != nil {
		Log.WriteString(err.Error() + "\n")
	} else {
		Log.WriteString("Enc was success")
	}
	CreateFileOpenFileWrite(filename, &network)
}
func dec (filename string){
	var prods []product.Keyboard
	n,err := ioutil.ReadFile(filename)
        if err != nil {
                fmt.Printf("cannot read file")
                panic(err)
        }
        p := bytes.NewBuffer(n)
        dec := gob.NewDecoder(p)
		err = dec.Decode(&prods)
        if err != nil {
			Log.WriteString(err.Error() + "\n")
        } else {
			Log.WriteString("Dec was success")
		}
	fmt.Println(prods)
}
var Filename = string(os.Args[1])

func main() {
	enc("Keyboards.bin")
	dec("Keyboards.bin")

}
