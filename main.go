package main 

import (
	"fmt"
	"net/http"
	"github.com/4Tune-Light/test-3/api/word"
)

func main() {
	http.Handle("/", word.WordAPI{})

	fmt.Print("Server is running on PORT: 3000")
	http.ListenAndServe(":3000", nil)
}