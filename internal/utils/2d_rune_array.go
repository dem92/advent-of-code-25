package utils

func Get2dRuneArray(content []string) [][]rune {
	runes := [][]rune{}

	for _, line := range content {
		runes = append(runes, []rune(line))
	}

	return runes
}
