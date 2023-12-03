func generateMatrix(n int) [][]int {
	x, y := 0, 0

	// Init matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	if n == 0 {
		return matrix
	}

	matrix[0][0] = 1
    	if n == 1 {
		return matrix
	}

    direction := 0
	prev := 1
	for direction != -1 {
		x, y, direction = move(x, y, direction, n, matrix)
		if x >= 0 && x < n && y >= 0 && y < n {
			prev++
			matrix[x][y] = prev
		}
	}

	return matrix
}

func move(x, y, direction, n int, matrix [][]int) (nextX, nextY, nextDirection int) {
	switch direction {
	case 0:
		if y == n-1 || y < n-1 && matrix[x][y+1] > 0 { // Reaching left most
			if x+1 < n && matrix[x+1][y] > 0 { //Next step is filled
				return -1, -1, -1
			}

			return x + 1, y, 1
		}

		// Keep moving to the right
		return x, y + 1, direction
	case 1:
		// Reaching the bottom
		if x == n-1 || x < n-1 && matrix[x+1][y] > 0 {
			if y-1 >= 0 && matrix[x][y-1] > 0 { //Next step is filled
				return -1, -1, -1
			}

			return x, y - 1, 2
		}

		// Keep moving down
		return x + 1, y, direction
	case 2:
		// Reaching the left most
		if y == 0 || y > 0 && matrix[x][y-1] > 0 {
			if x-1 >= 0 && matrix[x-1][y] > 0 { //Next step is filled
				return -1, -1, -1
			}

			return x - 1, y, 3
		}

		// Keep moving left
		return x, y - 1, direction

	case 3:
		// Reaching the top most
		if x == 0 || x > 0 && matrix[x-1][y] > 0 {
			if y+1 < n && matrix[x][y+1] > 0 { //Next step is filled
				return -1, -1, -1
			}

			return x, y + 1, 0
		}

		// Keep moving up
		return x - 1, y, direction
	}

	return -1, -1, -1
}