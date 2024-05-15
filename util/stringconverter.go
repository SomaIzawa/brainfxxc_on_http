package util

// 文字列の加工

import (
	"strings"
)

// 文字列から、特定の文字のみを抜き出します

func ExtractSpecificCharacters(str string, spchars []string) []string {
	chars := strings.Split(str, "")
	var modifiedStrings []string
	for _, c := range chars {
		if CompareStrings(c, spchars) {
			modifiedStrings = append(modifiedStrings, c)
		}
	}
	return modifiedStrings
}

// 1byteの文字を1byteの文字のリストと比較して、一致するものがあるかどうかを判定します

func CompareStrings(target string, compareList []string) bool {
	for _, compareItem := range compareList {
		if target == compareItem {
			return true
		}
	}
	return false
}