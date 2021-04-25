package sudoku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Success bool   `json:"response"`
	Size    string `json:"size"`
	Squares []Cell `json:"squares"`
}

func (r *Response) fromJson(data []byte) error {
	return json.Unmarshal(data, r)
}

func (b *Board) toJson() ([]byte, error) {
	return json.Marshal(b.Cells)
}

func fetch(size, difficulty uint8) []byte {
	url := fmt.Sprintf("http://www.cs.utep.edu/cheon/ws/sudoku/new/?size=%v&level=%v", size, difficulty)

	res, err := http.Get(url)
	must(err)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	must(err)
	return body
}
