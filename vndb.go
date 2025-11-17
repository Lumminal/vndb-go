package main

import (
	"fmt"
	"log"
	"os"
	"vndb-go/wrapper"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println(os.Getenv(wrapper.VNDBToken))
}
