package helper

func IsLost(round int, solutionWord string, guessWord string) bool {
	if round == 6 && solutionWord != guessWord {
		return true
	}
	if solutionWord == guessWord {
		return false
	}
	return true
}
