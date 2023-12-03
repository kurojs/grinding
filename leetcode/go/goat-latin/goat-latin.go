func toGoatLatin(S string) string {
    	vowels := map[string]bool{
		"a": true,
		"o": true,
		"e": true,
		"i": true,
		"u": true,
	}

	words := strings.Split(S, " ")
	goatLatins := []string{}
	postFix := "a"

	for _, word := range words {
		if vowels[strings.ToLower(string(word[0]))] {
			word += "ma"
		} else {
			word = word[1:] + string(word[0]) + "ma"
		}

		word += postFix
		goatLatins = append(goatLatins, word)
		postFix += "a"
	}

	return strings.Join(goatLatins, " ")
}