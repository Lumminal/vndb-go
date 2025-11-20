package tests

import (
	"log"
	"os"
	"testing"
	"vndb-go/vndb-go"
)

var clientTest *vndb_go.VNDBClient

func TestMain(M *testing.M) {
	token := os.Getenv(vndb_go.VNDBToken)
	if token == "" {
		log.Printf("No VNDB token found. Some tests may fail")
		os.Exit(M.Run())
		return
	}

	log.Printf("Starting tests...")
	clientTest = vndb_go.NewVndbClient(token)

	os.Exit(M.Run())
}
