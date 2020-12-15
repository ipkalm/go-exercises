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
	fmt.Println("assortion:", e.(*customErr).addInfo) // write (TYPE) correctly, it may be non-pointer like (customErr) or pointer (*customErr)
	log.Fatalln(e)
}
