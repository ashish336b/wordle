package helper

func InProgress(round int, solutionWord string, guessWord string) bool {
	if (round == 6 && solutionWord != guessWord) || solutionWord == guessWord {
		return false
	}

	return true
}

func IsWon(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord == guessWord
}
