
func isNumber(s string) bool {
  state := buildState(s)
  
  validState := map[string]bool{
    "2":  true,
    "12": true,
    "1232": true,
    "23": true,
    "32": true,
    "132": true,
    "123": true,
    "232": true,
    "242": true,
    "1242": true,
    "2412": true,
    "12412": true,
    "2342": true,
    "3242": true,
    "13242": true,
    "23242": true,
    "123242": true,
    "12342": true,
    "23412": true,
    "232412": true,
    "32412": true,
    "132412": true,
  }
  
  fmt.Println(state)
  
  return validState[state]
}


/*
State 1: "+" or "-"
State 2: numberic
State 3: "."
State 4: "e" or "E"
*/

func buildState(s string) string {
  state := ""
  for i := 0; i < len(s); i++ {
    if s[i] == '-' || s[i] == '+' {
      state += "1"
    } else if s[i] >= 48 && s[i] <= 57 { // 0-9
      state += "2"
      for i+1 < len(s) && s[i+1] >= 48 && s[i+1] <= 57 {
        i++
      }
    } else if s[i] == '.' {
      state += "3"
    } else if s[i] == 'e' || s[i] == 'E' {
      state += "4"
    } else {
      return ""
    }
  }  
  
  return state
}
