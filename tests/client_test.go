package tests

import (
	"log"
	"os"
	"testing"
	"vndb-go/wrapper"
)

var clientTest *wrapper.VNDBClient

func TestMain(M *testing.M) {
	token := os.Getenv(wrapper.VNDBToken)
	if token == "" {
		log.Printf("No VNDB token found. Some tests may fail")
		os.Exit(M.Run())
		return
	}

	log.Printf("Starting tests...")
	clientTest = wrapper.NewVndbClient(token)

	os.Exit(M.Run())
}
