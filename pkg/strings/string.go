package strings

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func ReplaceAtMap(in string, m map[int]rune) string {
	out := []rune(in)
	for k, v := range m {
		out[k] = v
	}
	return string(out)
}
