package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

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
	if !include(guess, Words) && !include(guess, ValidWord) && len(guess) == 5 {
		message = "Not in Word List"
	}
	if len(guess) != 5 {
		message = "words should be exactly 5"
	}
	return message, (include(guess, Words) || include(guess, ValidWord)) && len(guess) == 5
}
func CheckAccuracy(solution string, guessed string) string {
	solutionWordArr := strings.Split(solution, "")
	guessedWordArr := strings.Split(guessed, "")
	finalWord := ""
	for i, value := range solutionWordArr {
		if !include(guessedWordArr[i], solutionWordArr) {
			addToGrayWordList(guessedWordArr[i])
			finalWord += coloredString("grey").Render(guessedWordArr[i])
		}
		if guessedWordArr[i] == value {
			addToGreenWordList(guessedWordArr[i])
			solutionWordArr[i] = "*"
			finalWord += coloredString("green").
				Render(guessedWordArr[i])
		}
		if guessedWordArr[i] != value && include(guessedWordArr[i], solutionWordArr) {
			addToYellowWordList(guessedWordArr[i])
			solutionWordArr[index(guessedWordArr[i], solutionWordArr)] = "*"
			finalWord += coloredString("yellow").
				Render(guessedWordArr[i])
		}
	}
	return lipgloss.NewStyle().
		Margin(1).
		MarginLeft(5).
		Render(finalWord)
}
func GetRightAnswer(solution string) string {
	return coloredString("cyan").Margin(2).Render(solution)
}
func WinCondition(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord == guessWord
}
func LoseCondition(round int, solutionWord string, guessWord string) bool {
	return round == 6 && solutionWord != guessWord
}
func addToGreenWordList(s string) {
	// if alphabet is present in yellow word remove from yello word list
	if index(s, YellowWord) != -1 {
		YellowWord = RemoveIndex(index(s, YellowWord), YellowWord)
	}
	// if alphabet is not already persent in greenword list add to greenword list
	if index(s, GreenWord) == -1 {
		GreenWord = append(GreenWord, s)
	}
}
func addToYellowWordList(s string) {
	// already green word cannot be yellow again in next round
	// green word ko list ma xaena banee matra add garnee
	if index(s, GreenWord) == -1 {
		YellowWord = append(YellowWord, s)
	}
}
func addToGrayWordList(s string) {
	// word should not be in green and yello word list
	if !(index(s, GreenWord) != -1 || index(s, YellowWord) != -1) {
		GrayWord = append(GrayWord, s)
	}
}
func include(a string, list []string) bool {
	for _, b := range list {
		if strings.ToUpper(b) == strings.ToUpper(a) {
			return true
		}
	}
	return false
}
func index(a string, arr []string) int {
	for i, value := range arr {
		if value == a {
			return i
		}
	}
	return -1
}
func RemoveIndex(index int, s []string) []string {
	return append(s[:index], s[index+1:]...)
}
