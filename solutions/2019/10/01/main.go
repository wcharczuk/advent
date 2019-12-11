package main

func main() {

}

// Board is a 2 dimensional bool array.
type Board [][]bool

// AddRow adds a row to the board.
func (b *Board) AddRow(row ...bool) {
	*b = append(*b, row)
}

// Get returns a value at a given coordinate.
func (b Board) Get(x, y int) bool {
	if y < len(b) {
		if x < len(b[y]) {
			return b[y][x]
		}
		return b[y][x%len(b[y])]
	}
	if x < len(b[y%len(b)]) {
		return b[y%len(b)][x]
	}
	return b[y%len(b)][x%len(b[y%len(b)])]
}
