package products

import (
	/*"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"*/
)

type Product struct {
	Flag  string
	Name  string
	Type  string
	Rate  int64
	Price int64
}

type t interface{
	GetProduct() ([]Product)
	GetCount() (int)
	enc()()
	dec()()
}
