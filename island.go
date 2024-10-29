package livetest

func maxAreaOfIsland(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count = max(count, dfsIsland(grid, i, j))
			}
		}
	}

	return count
}

func dfsIsland(grid [][]int, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfsIsland(grid, i+1, j) + dfsIsland(grid, i-1, j) + dfsIsland(grid, i, j+1) + dfsIsland(grid, i, j-1)
}
