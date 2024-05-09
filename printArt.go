package art



func PrintArt(s string, a map[int][]string) string {
	if s == "" {
		return ""
	}

	var result string
	var asciiArt [][]string
	for _, p := range s {
		asciiArt = append(asciiArt, a[int(p)])
	}

	for i := 0; i < 8; i++ {
		for _, row := range asciiArt {
			result += row[i]
		}
		result += "\n"
	}

	return result
}