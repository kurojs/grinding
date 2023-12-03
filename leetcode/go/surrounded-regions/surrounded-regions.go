func solve(board [][]byte)  {
    visited := make(map[string]bool)
    m := len(board)
    if m == 0 {
        return
    }
    n := len(board[0])
    
    // Looking around the bound
    for x := 0; x < m; x++ {
        for y := 0; y < n; y++ {
            if board[x][y] == 'X' {
                continue
            }
            if x == 0 || x == m - 1 || y == 0 || y == n - 1 {
                expand(board, visited, x, y, m, n)
            }             
        }
    }
    
    for x := 0; x < m; x++ {
        for y := 0; y < n; y ++ {
            if visited[fmt.Sprintf("%d-%d", x, y)] || board[x][y] == 'X' {
                continue
            }
            
            board[x][y] = 'X'
        }
    }
}

func expand(board [][]byte, visited map[string]bool, x, y, m, n int) {
    if x < 0 || y < 0 || x >= m || y >= n {
        return
    }
    
    if visited[fmt.Sprintf("%d-%d", x, y)] {
        return    
    }
    
    if board[x][y] == 'X' {
        return
    }
    
    visited[fmt.Sprintf("%d-%d", x, y)] = true
    expand(board, visited, x + 1, y, m, n)
    expand(board, visited, x - 1, y, m, n)
    expand(board, visited, x, y + 1, m, n)
    expand(board, visited, x, y - 1, m, n)
}

