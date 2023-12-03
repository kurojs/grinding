func letterCombinations(digits string) []string {
    ans := make([]string, 0)
    if len(digits) == 0 {
        return ans
    }
    
    m := make(map[byte][]string)
    m['1'] = []string{}
    m['2'] = []string{"a", "b", "c"}
    m['3'] = []string{"d", "e", "f"}
    m['4'] = []string{"g", "h", "i"}
    m['5'] = []string{"j", "k", "l"}
    m['6'] = []string{"m", "n", "o"}
    m['7'] = []string{"p", "q", "r", "s"}
    m['8'] = []string{"t", "u", "v"}
    m['9'] = []string{"w", "x", "y", "z"}
    m['0'] = []string{" "}
    
    queue := make([][]string, 0)
    queue = append(queue, []string{""})
    idx := 0
    remain := 1
    
    for len(queue) > 0 {        
        if idx >= len(digits) {
            return ans
        }
        
        // Pop queue
        words := queue[0]
        queue = queue[1:]
        
        nextWords := make([]string, 0)
        letters := m[digits[idx]]
        for _, word := range words {
            for _, letter := range letters {
                newWord := word + letter
                nextWords = append(nextWords, newWord)
                if len(word + letter) == len(digits) {
                    ans = append(ans, newWord)
                }
            }
        }
        
        queue = append(queue, nextWords)
        remain--
        
        if remain == 0 {
            idx++
            remain = len(queue)
        }
    }
     
    return ans
}