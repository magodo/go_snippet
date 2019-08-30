package array

// Expand will expand the given 2d array, adding one col and row from start, and fill them with `fill`.
func Expand(grid [][]int, fill int) [][]int {
	nRow := len(grid) + 1
	nCol := len(grid[0]) + 1
	newGrid := make([][]int, nRow)
	for row := range newGrid {
		newGrid[row] = make([]int, nCol)
		newGrid[row][0] = fill
	}
	for idx := range newGrid[0] {
		newGrid[0][idx] = 0
	}

	for x := 0; x < nRow-1; x++ {
		for y := 0; y < nCol-1; y++ {
			newGrid[x+1][y+1] = grid[x][y]
		}
	}
	return newGrid
}
