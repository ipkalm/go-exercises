package main

import "fmt"

func main() {
	fmt.Println("Open file/connection do some stuff")
	defer close()
	fmt.Println("send information throught connection")
	fmt.Println("recieve some answer")
	fmt.Println("write answer into the file")
}

func close() {
	fmt.Println("I am defer func. Closing file/connection...")
}
