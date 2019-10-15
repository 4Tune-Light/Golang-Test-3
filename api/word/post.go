package word

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/4Tune-Light/test-3/controllers/word/test1"
)

type wordType struct {
    Word string
}

type wordRes struct {
	Vowels string
	Consonants string
  Result string
}

func DoPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
  var word wordType

  err := decoder.Decode(&word)
  if err != nil {
      fmt.Printf("Word is required")
      w.WriteHeader(http.StatusBadRequest)
      return
  }

  data := test1.Test1(word.Word)

  res := wordRes{data["vowels"], data["consonants"], data["result"]}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
}