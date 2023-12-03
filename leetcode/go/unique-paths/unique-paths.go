func uniquePaths(m int, n int) int {
    // We start from position [0:0]
    // End at [m:n]
    
    if m <= 0 || n <= 0 {
        return 0
    } 
    
    if m == 1 && n == 1 {
        return 1
    }
    
    // Steps will we store in [m][n]int
    var steps [][]int64
    for i := 0; i < m; i++ {
        steps = append(steps, make([]int64, n)) 
    }
    
    // Init state
    steps[m-1][n-1] = 1
    
    // Resolve main proplem by solving sub-problems
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            // Skip init value
            if i == m - 1 && j == n - 1 {
                continue
            }
            
            curr := int64(0)
            
            // currStep = leftStep + downStep
            if i + 1 < m {
                curr += steps[i + 1][j]
            }
            
            if j + 1 < n {
                curr += steps[i][j+1]
            }
            
            steps[i][j] = curr
        }
    }
    
    fmt.Println(steps)
    
    return int(steps[0][0])
}
