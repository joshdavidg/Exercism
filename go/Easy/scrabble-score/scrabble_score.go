// Package scrabble implements functions to calculate word scores
package scrabble

// Score returns the score of a passed in scrabble word
func Score(word string) int {
	score := 0
	for i := 0; i < len(word); i++ {
		score += scoreHelper(rune(word[i]))
	}

	return score
}

func scoreHelper(b rune) int {
	switch {
	case b == 65 || b == 69 || b == 73 || b == 79 || b == 85 || b == 76 || b == 78 || b == 82 || b == 83 || b == 84:
		fallthrough
	case b == 97 || b == 101 || b == 105 || b == 111 || b == 117 || b == 108 || b == 110 || b == 114 || b == 115 || b == 116:
		return 1
	case b == 68 || b == 71:
		fallthrough
	case b == 100 || b == 103:
		return 2
	case b == 66 || b == 67 || b == 77 || b == 80:
		fallthrough
	case b == 98 || b == 99 || b == 109 || b == 112:
		return 3
	case b == 70 || b == 72 || b == 86 || b == 87 || b == 89:
		fallthrough
	case b == 102 || b == 104 || b == 118 || b == 119 || b == 121:
		return 4
	case b == 75 || b == 107:
		return 5
	case b == 74 || b == 88:
		fallthrough
	case b == 106 || b == 120:
		return 8
	case b == 81 || b == 90:
		fallthrough
	case b == 113 || b == 122:
		return 10
	}

	return 0
}

// ScoreMultiplied returns the score for a word that has a letter multiplier
// which is determined by a map with the index of the letter as a key and the multiplier as the value
func LetterScoreMultiplier(word string, index map[int]int) int {
	score := 0
	for i := 0; i < len(word); i++ {
		if val, ok := index[i]; ok {
			score += scoreHelper(rune(word[i])) * val
			continue
		}
		score += scoreHelper(rune(word[i]))
	}

	return score
}

// WordScoreMultiplier returns the score of a word with a multiplier
func WordScoreMultiplier(word string, multiplier int) int {
	return Score(word) * multiplier
}
