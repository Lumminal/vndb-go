package main

import (
	"fmt"
	"os"
	"vndb-go/wrapper"
)

func main() {
	fmt.Println(os.Getenv(wrapper.VNDBToken))
}
