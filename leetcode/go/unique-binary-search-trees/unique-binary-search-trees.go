func numTrees(n int) int {
    dp := make([]int, n + 1)
    dp[0] = 1
    dp[1] = 1
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] += dp[j-1] * dp[i-j];
        }
    }
    
    return dp[n]
}

// 0 -> 1
// 1 -> 1
// 2 -> 2
// 3 -> 5 (+3)
// 4 -> 14 (+9)
// 5 -> 42 (+28)
// 6 -> 132
// 7 -> 429
// 8 -> 1430