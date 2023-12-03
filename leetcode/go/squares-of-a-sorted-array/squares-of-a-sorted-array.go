func sortedSquares(nums []int) []int {
    neg, pos := split(nums)
    return merge(neg, pos)
}

func split(nums []int) ([]int, []int) {
    neg, pos := make([]int, 0), make([]int, 0)
    for _, n := range nums {
        if n < 0 {
            neg = append(neg, n*n)
        } else {
            pos = append(pos, n*n)
        }
    }
    
    return neg, pos
}

func merge(neg, pos []int) []int {
    i, j := len(neg)-1, 0
    lpos := len(pos)
    
    merged := make([]int, 0, len(neg)+lpos)
    for i >= 0 && j < lpos {
        if neg[i] < pos[j] {
            merged = append(merged, neg[i])
            i--
            continue
        }
        
        merged = append(merged, pos[j])
        j++
    }
    
    for i >= 0 {
        merged = append(merged, neg[i])
        i--
    }
    
    for j < lpos {
        merged = append(merged, pos[j])
        j++
    }
    
    return merged
}