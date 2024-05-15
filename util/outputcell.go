package util

// コマンドラインへの出力補助ツール

import (
	"fmt"
	"strconv"
	"strings"
)

// 罫線「+---+」を引く

func OutPutLine(width int){
	for i:=0; i<width; i++ {
		fmt.Print("+---")
	}
	fmt.Print("+\n")
}

// 何もない空行

func OutPutEmptyLine(){
	fmt.Print("\n")
}

// 配列をセルにして出力

func OutPutValues(array []int){
	for _, item := range array {
		fmt.Print("|")
		fmt.Print(OutPutIntOnXByte(item, 3))
	}
	fmt.Print("|\n")
}

// 指定されたbyte数で右詰にして文字列を返却（ 例：3byteで "2" は "  2" ）

func OutPutIntOnXByte(input interface{}, byte int) string{
	var inputStr string

	switch v := input.(type) {
	case int:
		inputStr = strconv.Itoa(v)
	case string:
		inputStr = v
	default:
		return ""
	}

	numSpaces := byte - len(inputStr)
	if numSpaces < 0 {
		numSpaces = 0
	}

	outputStr := strings.Repeat(" ", numSpaces) + inputStr
	return outputStr
}