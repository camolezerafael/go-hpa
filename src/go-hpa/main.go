package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)
func greeting( text string ) string{
	x := 0.0001

	for i := 0; i < 1000000000; i++{
		x += math.Sqrt(x)
	}

	return fmt.Sprintf("<b>%s</b>", text)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}