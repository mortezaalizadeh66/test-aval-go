package main

import(
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("Go Example")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
