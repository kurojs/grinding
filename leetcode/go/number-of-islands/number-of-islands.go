func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	visited := make(map[string]bool)
	islands := 0

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if !visited[fmt.Sprint("%d-%d", x, y)] && grid[x][y] == '1' {
				expand(grid, visited, x, y, m, n)
				islands++
			}
		}
	}

	return islands
}

func expand(grid [][]byte, visited map[string]bool, x, y, m, n int) {
	if x >= m || x < 0 || y >= n || y < 0 {
		return
	}

	if grid[x][y] == '0' {
		return
	}

	if visited[fmt.Sprint("%d-%d", x, y)] {
		return
	}

	visited[fmt.Sprint("%d-%d", x, y)] = true

	expand(grid, visited, x+1, y, m, n)
	expand(grid, visited, x-1, y, m, n)
	expand(grid, visited, x, y+1, m, n)
	expand(grid, visited, x, y-1, m, n)
}