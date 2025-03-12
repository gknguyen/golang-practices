package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("failed to read balance file")
	}

	balanceText := string(data)

	balance, err2 := strconv.ParseFloat(balanceText, 64)
	if err2 != nil {
		return 0, errors.New("failed to parse balance")
	}

	return balance, nil
}
