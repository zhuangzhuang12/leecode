package day1

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfsLands(grid, i, j)
			}
		}
	}

	return res
}

func dfsLands(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
	}
	dfsLands(grid, i-1, j)
	dfsLands(grid, i+1, j)
	dfsLands(grid, i, j-1)
	dfsLands(grid, i, j+1)
}
