package main

import (
	"fmt"
	"log"
)

type customErr struct {
	addInfo string
}

func (c *customErr) Error() string {
	return fmt.Sprintf("add info :: %v", c.addInfo)
}

func main() {
	ce := customErr{"something go wrong!"}
	foo(&ce)
}

func foo(e error) {
	log.Fatalln(e)
}
