package main

import (
	"os"
	"fmt"
)

func main() {
	path, _ := os.Executable()
	fmt.Println(path)
}
