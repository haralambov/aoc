package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/haralambov/aoc/lib/input"
)

const (
	FiveOfAKindPower  = 7
	FourOfAKindPower  = 6
	FullHousePower    = 5
	ThreeOfAKindPower = 4
	TwoPairsPower     = 3
	OnePairPower      = 2
	HighCardPower     = 1
)

func solve(isPartTwo bool) string {
	inputLines := input.GetInputAsSlice(2023, 7)
	// hand -> bid
	inputData := make(map[string]int, len(inputLines))
	// calculated hand power -> list of hands e.g.:
	// Five of a kind: power 7 -> [list of five of a kind hands]
	// Full houses   : power 5 -> [list of full house hands]
	handPowers := make(map[int][]string, len(inputLines))
	// calculated hand high card power: hand -> power
	highCardPowers := make(map[string]int, len(inputLines))
	highCardPowersReversed := make(map[int]string, len(inputLines))

	for _, inputLine := range inputLines {
		if trimmedLine := strings.TrimSpace(inputLine); len(trimmedLine) > 0 {
			trimmedLineParts := strings.Split(trimmedLine, " ")
			hand := trimmedLineParts[0]
			bet, _ := strconv.Atoi(trimmedLineParts[1])
			inputData[hand] = bet

			handPower := calcHandPower(hand, isPartTwo)
			handPowers[handPower] = append(handPowers[handPower], hand)

			highCardPower := calcHighCardPower(hand, isPartTwo)
			highCardPowerConverted, _ := strconv.ParseInt(highCardPower, 16, 32)
			highCardPowers[hand] = int(highCardPowerConverted)
			highCardPowersReversed[int(highCardPowerConverted)] = hand
		}
	}

	orderedHands := make([]string, 0)
	for _, handPower := range getAllHandPowers() {
		if _, ok := handPowers[handPower]; ok {
			handsWithThatPower := handPowers[handPower]
			handsHighCardPowers := make([]int, 0)
			for _, hand := range handsWithThatPower {
				handsHighCardPowers = append(handsHighCardPowers, highCardPowers[hand])
			}
			sort.Sort(sort.Reverse(sort.IntSlice(handsHighCardPowers)))
			for _, highCardPower := range handsHighCardPowers {
				hand := highCardPowersReversed[highCardPower]
				orderedHands = append(orderedHands, hand)
			}
		}
	}

	totalWinnings := 0
	orderedHandsSize := len(orderedHands)
	for handIndex, hand := range orderedHands {
		bet := inputData[hand]
		winning := (orderedHandsSize - handIndex) * bet
		totalWinnings += winning
	}
	return fmt.Sprintf("%v", totalWinnings)
}

// Each card in the hand is represented as a hex string
// which is later on converted to int for easier comparison
func calcHighCardPower(hand string, isPartTwo bool) string {
	cardPower := ""
	for _, char := range hand {
		cardPower += getCardPower(char, isPartTwo)
	}
	return cardPower
}

func getCardPower(char rune, isPartTwo bool) string {
	switch char {
	case 'A':
		return "E"
	case 'K':
		return "D"
	case 'Q':
		return "C"
	case 'J':
		// In part two Jacks are Jokers,
		// and are worth less then the 2s
		if isPartTwo {
			return "1"
		}
		return "B"
	case 'T':
		return "A"
	case '9':
		return "9"
	case '8':
		return "8"
	case '7':
		return "7"
	case '6':
		return "6"
	case '5':
		return "5"
	case '4':
		return "4"
	case '3':
		return "3"
	default:
		return "2"
	}
}

func calcHandPower(hand string, isPartTwo bool) int {
	jokersCount := getJokersCount(hand, isPartTwo)
	if isFiveOfAKind(hand) {
		return FiveOfAKindPower
	} else if isFourOfAKind(hand) {
		if isPartTwo && jokersCount > 0 { // Whether it's JKKKK or JJJJK, it's 5 of a kind
			return FiveOfAKindPower
		}
		return FourOfAKindPower
	} else if isFullHouse(hand) {
		if isPartTwo && jokersCount > 0 { // Whether it's JJKKK or JJJKK, it's 5 of a kind
			return FiveOfAKindPower
		}
		return FullHousePower
	} else if isThreeOfAKind(hand) {
		if isPartTwo && jokersCount > 0 { // Whether it's JJJKQ or KKKQJ, it's 4 of a kind
			return FourOfAKindPower
		}
		return ThreeOfAKindPower
	} else if isTwoPairs(hand) {
		if isPartTwo {
			if jokersCount == 1 { // AAKKJ => full house
				return FullHousePower
			} else if jokersCount == 2 { // JJKKQ => 4 of a kind
				return FourOfAKindPower
			}
		}
		return TwoPairsPower
	} else if isOnePair(hand) {
		if isPartTwo && jokersCount > 0 { // Whether it's AAJKQ or JJAKQ, it's 3 of a kind
			return ThreeOfAKindPower
		}
		return OnePairPower
	}
	if isPartTwo && jokersCount > 0 { // AKQTJ => AKQTT - always 1 pair
		return OnePairPower
	}
	return HighCardPower
}

func getJokersCount(hand string, isPartTwo bool) int {
	if !isPartTwo {
		return 0
	}
	return strings.Count(hand, "J")
}

func getAllHandPowers() []int {
	return []int{
		FiveOfAKindPower,
		FourOfAKindPower,
		FullHousePower,
		ThreeOfAKindPower,
		TwoPairsPower,
		OnePairPower,
		HighCardPower,
	}
}

func isFiveOfAKind(hand string) bool {
	chars := make(map[rune]int, 0)
	for _, char := range hand {
		if _, ok := chars[char]; ok {
			chars[char] = chars[char] + 1
		} else {
			chars[char] = 1
		}
	}
	return len(chars) == 1
}

func isFourOfAKind(hand string) bool {
	chars := make(map[rune]int, 0)
	keys := make([]rune, 0)
	for _, char := range hand {
		if _, ok := chars[char]; ok {
			chars[char] = chars[char] + 1
		} else {
			keys = append(keys, char)
			chars[char] = 1
		}
	}
	if len(chars) != 2 {
		return false
	}
	firstCardTypeCount := chars[keys[0]]
	secondCardTypeCount := chars[keys[1]]
	isFullHouse := firstCardTypeCount*secondCardTypeCount == 4
	return isFullHouse
}

func isFullHouse(hand string) bool {
	chars := make(map[rune]int, 0)
	keys := make([]rune, 0)
	for _, char := range hand {
		if _, ok := chars[char]; ok {
			chars[char] = chars[char] + 1
		} else {
			keys = append(keys, char)
			chars[char] = 1
		}
	}
	if len(chars) != 2 {
		return false
	}
	firstCardTypeCount := chars[keys[0]]
	secondCardTypeCount := chars[keys[1]]
	isFullHouse := firstCardTypeCount*secondCardTypeCount == 6
	return isFullHouse
}

func isThreeOfAKind(hand string) bool {
	chars := make(map[rune]int, 0)
	keys := make([]rune, 0)
	for _, char := range hand {
		if _, ok := chars[char]; ok {
			chars[char] = chars[char] + 1
		} else {
			keys = append(keys, char)
			chars[char] = 1
		}
	}
	if len(chars) != 3 {
		return false
	}
	firstCardTypeCount := chars[keys[0]]
	secondCardTypeCount := chars[keys[1]]
	thirdCardTypeCount := chars[keys[2]]
	isThreeOfAKind := firstCardTypeCount*secondCardTypeCount*thirdCardTypeCount == 3
	return isThreeOfAKind
}

func isTwoPairs(hand string) bool {
	chars := make(map[rune]int, 0)
	keys := make([]rune, 0)
	for _, char := range hand {
		if _, ok := chars[char]; ok {
			chars[char] = chars[char] + 1
		} else {
			keys = append(keys, char)
			chars[char] = 1
		}
	}
	if len(chars) != 3 {
		return false
	}
	firstCardTypeCount := chars[keys[0]]
	secondCardTypeCount := chars[keys[1]]
	thirdCardTypeCount := chars[keys[2]]
	isThreeOfAKind := firstCardTypeCount*secondCardTypeCount*thirdCardTypeCount == 4
	return isThreeOfAKind
}

func isOnePair(hand string) bool {
	chars := make(map[rune]int, 0)
	keys := make([]rune, 0)
	for _, char := range hand {
		if _, ok := chars[char]; ok {
			chars[char] = chars[char] + 1
		} else {
			keys = append(keys, char)
			chars[char] = 1
		}
	}
	if len(chars) != 4 {
		return false
	}
	firstCardTypeCount := chars[keys[0]]
	secondCardTypeCount := chars[keys[1]]
	thirdCardTypeCount := chars[keys[2]]
	fourthCardTypeCount := chars[keys[3]]
	isThreeOfAKind := firstCardTypeCount*secondCardTypeCount*thirdCardTypeCount*fourthCardTypeCount == 2
	return isThreeOfAKind
}
