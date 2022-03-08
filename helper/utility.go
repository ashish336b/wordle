package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func GetSolution() string {
	rand.Seed(time.Now().UnixNano())
	solutionIndex := rand.Intn(len(Words))
	solutionWord := strings.ToUpper(Words[solutionIndex])
	return solutionWord
}
func ValidateWord(guess string) (string, bool) {
	message := ""
	if !include(guess, Words) && len(guess) == 5 {
		message = "Not in Word List"
	}
	if len(guess) != 5 {
		message = "words should be exactly 5"
	}
	return message, include(guess, Words) && len(guess) == 5
}
func CheckAccuracy(solution string, guessed string) string {
	solutionWordArr := strings.Split(solution, "")
	guessedWordArr := strings.Split(guessed, "")
	finalWord := ""
	for i, value := range solutionWordArr {
		if guessedWordArr[i] == value {
			finalWord += baseStyle().
				Foreground(lipgloss.Color("#14c700")).
				Render(guessedWordArr[i])
		}
		if guessedWordArr[i] != value && include(guessedWordArr[i], solutionWordArr) {
			finalWord += baseStyle().
				Foreground(lipgloss.Color("#e6d600")).
				Render(guessedWordArr[i])
		}
		if !include(guessedWordArr[i], solutionWordArr) {
			finalWord += baseStyle().Render(guessedWordArr[i])
		}
	}
	return lipgloss.NewStyle().
		Margin(1).
		MarginLeft(5).
		Render(finalWord)
}

// use pointer concept
func GetRightAnswer(solution string) string {
	return baseStyle().Foreground(lipgloss.Color("#42C2FF")).Margin(2).Render(solution)
}
func InProgress(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord != guessWord
}

func WinCondition(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord == guessWord
}

func LoseCondition(round int, solutionWord string, guessWord string) bool {
	// fmt.Print(round == 6 && solutionWord != guessWord) //remove
	return round == 6 && solutionWord != guessWord
}
func include(a string, list []string) bool {
	for _, b := range list {
		if strings.ToUpper(b) == strings.ToUpper(a) {
			return true
		}
	}
	return false
}
