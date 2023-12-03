func countSegments(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    seg := 0
    for i := 0; i < len(s);  {
        if s[i] != ' ' {
            for i < len(s) && s[i] != ' ' {
                i++
            }
            
            seg++
            continue
        }
        
        for i < len(s) && s[i] == ' ' {
            i++
        }
    }
    
    return seg
}