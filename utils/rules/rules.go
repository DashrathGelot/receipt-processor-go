package rules

import (
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
)

func ConverToFloat(num string) (float64, error) {
	return strconv.ParseFloat(num, 64)
}

func Alphanumeric(retailer string) int {
	return len(regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(retailer, ""))
}

func IsTotalRound(total float64) bool {
	return math.Round(total) == total
}

func IsTotalMultiple(total float64) bool {
	return math.Mod(total, 0.25) == 0
}

func Rule4Items(items []models.Item) int {
	return (len(items) / 2) * 5
}

func Rule5TrimDesc(items []models.Item) (int, error) {
	points := 0
	for _, item := range items {
		descSize := len(strings.TrimSpace(item.ShortDescription))
		if descSize%3 == 0 {
			if price, err := ConverToFloat(item.Price); err == nil {
				points += int(math.Ceil(price * 0.2))
			} else {
				return int(price), err
			}
		}
	}
	return points, nil
}

func IsPurchaseDateOdd(purchaseDate string) (bool, error) {
	day, err := strconv.Atoi(strings.Split(purchaseDate, "-")[2])
	if err == nil {
		return day%2 == 1, nil
	}
	return false, err
}

func IsTime2to4(purchaseTime string) (bool, error) {
	hour, err := strconv.Atoi(strings.Split(purchaseTime, ":")[0])
	if err == nil {
		return hour >= 14 && hour <= 16, nil
	}
	return false, err
}
