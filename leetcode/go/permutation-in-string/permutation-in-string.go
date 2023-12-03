func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    if len(s1) == 0 {
        return true
    }
    
    need := make(map[byte]int)
    for i := range s1 {
        need[s1[i]]++
    }
    
    have := make(map[byte]int)
    for i := 0; i < len(s1); i++ {
        have[s2[i]]++
    }
    
    if isMatch(need, have) == 0 {
        return true
    }
    
    for i := len(s1); i < len(s2); i++ {
        have[s2[i-len(s1)]]--
        have[s2[i]]++
        
        if isMatch(need, have) == 0 {
            return true
        }
    }
    
    return false
}

func isMatch(need, have map[byte]int) int {
    for k, v := range need {
        if have[k] > v {
            return 1
        } else if have[k] < v {
            return -1
        }
    }
    
    return 0
}