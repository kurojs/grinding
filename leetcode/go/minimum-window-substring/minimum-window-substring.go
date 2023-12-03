func minWindow(s string, t string) string {
    // Initial require characers
    require := 0
    mRequired := make(map[byte]int)
    for i, _ := range t {
        mRequired[t[i]]++
        require++
    }
    
    // Using twos pointer left and right
    left, right := 0, 0
    minStart, minLen := -1, int(^uint(0)>>1)
    match := 0
    
    window := make(map[byte]int)
    for left <= right && right < len(s) {
        char := s[right]
        window[char]++
        
        // char exists in t
        if window[char] <= mRequired[char] {
            match++
        }
        
        for match == require {
            // Validate if current window has minimium len
            if currentLen := right - left + 1; minLen > currentLen {
                minLen = currentLen
                minStart = left
            }
            
            removedChar := s[left]
            left++
            window[removedChar]--
            if mRequired[removedChar] > 0 && window[removedChar] < mRequired[removedChar] {
                match--
            }
        }
        
        right++
    }
    
    if minStart == -1 {
        return ""
    }
    
    return s[minStart:minStart+minLen]
}