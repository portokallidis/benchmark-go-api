package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// The `json:"whatever"` bit is a way to tell the JSON
// encoder and decoder to use those names instead of the
// capitalised names
type item struct {
        Date string `json:"date"`
        Name string `json:"name"`
        Address string `json:"address"`
}
var sample *item = &item{
	Date: "2019-12-12",
	Name: "Nick Porto",
	Address: "No. 189, Grove St, Los Angeles",
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {

        switch r.Method {
        case "GET":
                // Just send out the JSON version of 'sample'
				j, _ := json.Marshal(sample)
				time.Sleep(1000 * time.Millisecond)
                w.Write(j)
        case "POST":
                // Decode the JSON in the body and overwrite 'sample' with it
                d := json.NewDecoder(r.Body)
                p := &item{}
                err := d.Decode(p)
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                }
                sample = p
        default:
                w.WriteHeader(http.StatusMethodNotAllowed)
                fmt.Fprintf(w, "This method is not supported")
        }
}

func main() {
        http.HandleFunc("/", sampleHandler)

        log.Println("Go!")
        http.ListenAndServe(":8080", nil)
}