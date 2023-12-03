func isHappy(n int) bool {
    return helper(n, make(map[int]bool))
}

func helper(n int, checked map[int]bool) bool {
    if n == 1 {
        return true
    }
    
    if checked[n] {
        return false
    }
    checked[n]= true
    
    sum := 0
    for n > 0 {
        tmp := n % 10
        sum += tmp * tmp
        n = n/10
    }
    
    return helper(sum, checked)
}
