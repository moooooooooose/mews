package main

import (
	"flag"
	"fmt"
	"github.com/moooooooooose/mews/pkg/sheetsapi"
	"net/http"
	"os"
)

var token string
var sheetID string

func main() {
	flag.StringVar(&token, "token", "", "google auth token")
	flag.StringVar(&sheetID, "sheet", "", "google sheet id")
	flag.Parse()

	if token == "" || sheetID == "" {
		flag.Usage()
		os.Exit(1)
	}

	sheetsClient := sheetsapi.NewClient(http.DefaultClient)
	_, err := sheetsClient.Get(sheetID, sheetsapi.RequestOptions{AuthToken: token})
	if err != nil {
		panic(err)
	}
	fmt.Println("i found your spreadsheet lol jk i hacked you, get rekt nerd")
}
