func maxAreaOfIsland(grid [][]int) int {
    max := 0
    n := len(grid)
    m := len(grid[0])
    
    queue := make([][2]int, 0)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                area := 0
                queue = append(queue, [2]int{i, j})
                
                for len(queue) > 0 {
                    idx := queue[0]
                    queue = queue[1:]
                    
                    
                    if idx[0] < 0 || idx[1] < 0 || idx[0] >= n || idx[1] >= m {
                        continue
                    }
                    if grid[idx[0]][idx[1]] != 1 {
                        continue
                    }
                    
                    area++
                    
                    for _, d := range directions {
                        queue = append(queue, [2]int{
                            idx[0]+d[0],
                            idx[1]+d[1],
                        })                               
                    }
                    grid[idx[0]][idx[1]] = -1
                }
                
                if area > max {
                    max = area
                }
            }
        }
    }
    
    return max
}

var directions = [][]int{
    {1, 0},
    {-1, 0},
    {0, 1},
    {0, -1},
}