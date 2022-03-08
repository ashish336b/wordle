package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var MAX_CHANCES = 6

func RandomSolution() string {
	rand.Seed(time.Now().UnixNano())
	solutionIndex := rand.Intn(len(Words))
	solutionWord := strings.ToUpper(Words[solutionIndex])
	return solutionWord
}
func OfficialSolution() string {
	// Official intital date for wordle is 2021/06/19
	// so calculate no of days from initial date to today date
	// modulus of total days to number of array is index of official word.
	firstWordleDay := time.Date(2021, time.Month(6), 19, 0, 0, 0, 0, time.UTC) // 2021/06/19
	day := int(time.Now().Sub(firstWordleDay).Hours() / 24)                    // 2021/06/19 minus today_date
	officialWord := Words[day%len(Words)]
	return strings.ToUpper(officialWord)
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
			finalWord += coloredString("green").
				Render(guessedWordArr[i])
		}
		if guessedWordArr[i] != value && include(guessedWordArr[i], solutionWordArr) {
			finalWord += coloredString("yellow").
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
	return coloredString("cyan").Margin(2).Render(solution)
}

// todo: remove this logic
func InProgress(round int, solutionWord string, guessWord string) bool {
	return round <= MAX_CHANCES && solutionWord != guessWord
}

func WinCondition(round int, solutionWord string, guessWord string) bool {
	return round <= MAX_CHANCES && solutionWord == guessWord
}

func LoseCondition(round int, solutionWord string, guessWord string) bool {
	return round == MAX_CHANCES && solutionWord != guessWord
}
func include(a string, list []string) bool {
	for _, b := range list {
		if strings.ToUpper(b) == strings.ToUpper(a) {
			return true
		}
	}
	return false
}
