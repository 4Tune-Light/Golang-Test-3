package main 

import (
	"fmt"
	"net/http"
	"./api/word"
)

func main() {
	http.Handle("/", word.WordAPI{})

	fmt.Print("Server is running on PORT: 3000")
	http.ListenAndServe(":3000", nil)
}