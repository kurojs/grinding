func letterCasePermutation(s string) []string {
    if len(s) == 0 {
        return []string{}
    }
    
    mChar := make(map[byte]bool)
    mChar[toLower(s[0])] = true
    mChar[toUpper(s[0])] = true
    
    perms := make([]string, 0)
    if len(s) == 1{
        for c := range mChar {
            perms = append(perms, string(c))
        }
        
        return perms
    }
    
    nextPerms := letterCasePermutation(s[1:])
    for c := range mChar {
        for _, perm := range nextPerms {
            perms = append(perms, string(c) + perm)
        }
    }
    
    return perms
}

func toLower(c byte) byte {
    if c >= 65 && c <= 90 {
        return c + 32
    }
    
    return c
}

func toUpper(c byte) byte {
    if c >= 97 && c <= 122 {
        return c - 32
    }
    
    return c
}