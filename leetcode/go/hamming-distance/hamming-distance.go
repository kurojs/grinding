// Can using XOR for easy implementation but Go didn't it to me 
func hammingDistance(x int, y int) int {
    ans := 0
    for i := 0; i < 32; i++ {
        xbit := 1 & x
        ybit := 1 & y
        
        if xbit != ybit {
            ans++
        }
        
        x = x >> 1
        y = y >> 1
    }
    
    return ans
}


