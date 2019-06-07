package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	f "fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var urlmap = make(map[string]string)

func fillMap() {
	filePtr := flag.String(
		"csv",
		"shorturls.csv",
		"a csv file with the format shortened-url,long-url")

	flag.Parse()

	csvFile, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		urlmap[line[0]] = line[1]
	}

}

func redirect(w http.ResponseWriter, r *http.Request) {
	link := strings.TrimLeft(r.URL.Path, "/")
	f.Println(link)

	if val, ok := urlmap[link]; ok {
		http.Redirect(w, r, val, 308)
	}
}

func main() {
	fillMap()
	http.HandleFunc("/", redirect)
	f.Println("Server running on port 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
