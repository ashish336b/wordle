package helper

func InProgress(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord != guessWord
}

func IsWon(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord == guessWord
}

func IsLost(round int, solutionWord string, guessWord string) bool {
	return round == 6 && solutionWord != guessWord
}
