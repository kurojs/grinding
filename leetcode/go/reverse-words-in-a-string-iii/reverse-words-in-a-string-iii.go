func reverseWords(s string) string {
    result := ""
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            result += " "
        } else {
            word := ""
            for i < len(s) && s[i] != ' ' {
                word += string(s[i])
                i++
            }
            
            for j := range word {
                result += string(word[len(word)-1-j])
            }
            
            i--
        }
    }

    return result
}


/*
    [0, 1, 2, 3, 4, 5, 6]
    [a, b, _, d, e, f, g]
*/