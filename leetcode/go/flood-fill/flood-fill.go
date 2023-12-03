func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    n := len(image)
    m := len(image[0])
    
    visited := make([][]int, n)
    for i := range visited {
        visited[i] = make([]int, m)
    }
    
    queue := make([][2]int, 0)
    queue = append(queue, [2]int{sr, sc})
    color := image[sr][sc]
    
    for len(queue) > 0 {        
        p := queue[0]
        queue = queue[1:]
        
        if p[0] < 0 || p[0] >= n || p[1] < 0 || p[1] >= m {
            continue
        }
        
        if visited[p[0]][p[1]] == 1 || image[p[0]][p[1]] != color {
            continue
        }
        
        visited[p[0]][p[1]] = 1
        image[p[0]][p[1]] = newColor
        for _, direction := range directions {
            queue = append(queue, [2]int{p[0]+direction[0], p[1]+direction[1]})
        }
    }
    
    return image
}

var directions = [4][2]int{
    {1, 0},
    {-1, 0},
    {0, -1},
    {0, 1},
}