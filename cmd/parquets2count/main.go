package main

import (
	"fmt"
	"log"

	pc "github.com/takanoriyanagitani/go-parquets2count"
)

func main() {
	cnt, e := pc.StdinToFilenamesToParquetsToTotalCount()
	if nil != e {
		log.Fatal(e)
	}

	fmt.Printf("%v\n", cnt) //nolint:forbidigo
}
