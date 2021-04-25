package sudoku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Board struct {
	Size       uint8
	Difficulty uint8
	Cells      []Cell `json:"squares"`
}

type Response struct {
	Response bool   `json:"response"`
	Size     string `json:"size"`
	Squares  []Cell `json:"squares"`
}

type Cell struct {
	X   uint8 `json:"x"`
	Y   uint8 `json:"y"`
	Val uint8 `json:"value"`
}

func CreateBoard(size, difficulty uint8) *Board {
	return &Board{Size: size, Difficulty: difficulty, Cells: []Cell{}}
}

func (b *Board) Fetch() []byte {
	url := fmt.Sprintf("http://www.cs.utep.edu/cheon/ws/sudoku/new/?size=%v&level=%v", b.Size, b.Difficulty)

	res, err := http.Get(url)
	Must(err)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	Must(err)
	return body
}

func (b *Board) Print() {
	fmt.Printf("Size: %v\nDifficulty: %v\n\n", b.Size, b.Difficulty)
	for i := range b.Cells {
		c := b.Cells[i]
		fmt.Printf("{ X: %v, Y: %v, Value: %v }\n", c.X, c.Y, c.Val)
	}
}

func (r *Response) FromJson(data []byte) error {
	return json.Unmarshal(data, r)
}

// func (g *Game) ToJson() ([]byte, error) {
// 	return json.Marshal(&g.Cells)
// }

func Must(err error) {
	if err != nil {
		log.Fatalf("Error, %v\n", err)
	}
}
