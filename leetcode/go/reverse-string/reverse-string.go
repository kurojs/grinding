func reverseString(s []byte)  {
    mid := len(s) / 2
    end := len(s) -1
    for i := 0; i < mid; i++ {
        t := s[i]
        s[i] = s[end - i]
        s[end - i] = t
    }
}

/*
odd length: 1; 3; 5 7: 0 - n /2
even: 2, 4, 6: 
*/