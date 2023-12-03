func addStrings(num1 string, num2 string) string {
    if len(num1) < len(num2) {
        return addStrings(num2, num1)
    }
    
    remain := 0
    idx1 := len(num1) - 1
    idx2 := len(num2) - 1
    ans := ""
    
    for idx1 >= 0 && idx2 >= 0 {
        s := remain + toInt(num1[idx1]) + toInt(num2[idx2])
        remain = s / 10
        s = s % 10
        
        ans = fmt.Sprintf("%d", s) + ans

        idx1--
        idx2--
    }
    
    for idx1 >= 0 {
        s := remain + toInt(num1[idx1]) 
        remain = s / 10
        s = s % 10
        ans = fmt.Sprintf("%d", s) + ans
        idx1--
    }
    
    if remain > 0 {
        ans = fmt.Sprintf("%d", remain) + ans
    }
    
    return ans
}


func toInt(b byte) int {
    n, _ := strconv.Atoi(string(b))
    return n
}