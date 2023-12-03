func stoneGame(piles []int) bool {
    stoneMap := make([]int, 501, 501)
    for _, stone := range piles {
        stoneMap[stone]++
    }
    
    alex, lee := 0, 0
    for i := 500; i > 0; i-- {
        if stoneMap[i] == 0 {
            continue
        }
        
        isOdd := stoneMap[i] % 2
        alexTake := stoneMap[i] / 2 + isOdd
        alex += i * alexTake
        lee += i * (stoneMap[i] - alexTake)
    }
    
    return alex > lee
}