package sudoku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DTO for mapping from api
type Response struct {
	Success bool   `json:"response"`
	Size    string `json:"size"`
	Squares []Cell `json:"squares"`
}

// Unmarshalls response into a data transfer object
func mapData(b *Board) *Response {
	body := getBody(b.Size, b.Difficulty)
	dto := new(Response)
	must(dto.fromJson(body))
	return dto
}

// Filters only the body from a data response
func getBody(size, difficulty int) []byte {
	res := getData(size, difficulty)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	must(err)
	return body
}

// Fetches a randomized sudoku board from a public api
func getData(size, difficulty int) *http.Response {
	url := fmt.Sprintf("http://www.cs.utep.edu/cheon/ws/sudoku/new/?size=%v&level=%v", size, difficulty)
	res, err := http.Get(url)
	must(err)
	return res
}

// Turns json into a valid go object
func (r *Response) fromJson(data []byte) error {
	return json.Unmarshal(data, r)
}

// Tuns a valid go object into json
func (b *Board) toJson() ([]byte, error) {
	return json.Marshal(b.Cells)
}
