func countSubstrings(s string) int {
  palindromic := 0
  for i := range s {
    palindromic += expand(s, i, i)
    palindromic += expand(s, i, i+1)
  }
  
  return palindromic
}

func expand(s string, left, right int) int {
  palindromic := 0
  for left >= 0 && right < len(s) && s[left] == s[right] {
    left--
    right++
    palindromic++
  }
  
  return palindromic
}
