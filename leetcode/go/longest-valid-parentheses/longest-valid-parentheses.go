func longestValidParentheses(s string) int {
    dp := make([]int, len(s)) // max valid length at index i
    max := 0
    
    for i := 1; i < len(s); i++ {
        if s[i] == '(' {
            continue
        }
        
        // case ...()
        if s[i] == ')' && s[i-1] =='(' {
            if i == 1 {
                dp[i] = 2
            } else {
                dp[i] = dp[i-2] + 2
            }
        }
        
        // case ...))
        if s[i] == ')' && s[i-1] == ')' && (i - dp[i-1] - 1) >= 0 && s[i - dp[i-1] - 1] == '('{
            // 0 1 2 3 4 5 6 7 8
            // ( ( ( ) ( ) ) ) (
            // 0 0 0 2 0 4 6 8 0
            if i - dp[i-1] - 2 < 0 {
                dp[i] = dp[i-1] + 2
            } else {
                dp[i] = dp[i-1] + dp[i - dp[i-1] - 2] + 2
            }
        }
        
        if max < dp[i] {
            max = dp[i]
        }
    }
    
    return max
}

// longest substr -> sliding window

// valid prenthese -> stack
