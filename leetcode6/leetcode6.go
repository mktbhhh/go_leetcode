package main

import "fmt"

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""

	str := make([][]byte, numRows)

	index := 0

	for ; index < len(s); index++ {
		if index%(2*numRows-2) <= numRows-1 {
			str[index%(2*numRows-2)] = append(str[index%(2*numRows-2)], s[index])
		} else {
			str[2*numRows-index%(2*numRows-2)-2] = append(str[2*numRows-index%(2*numRows-2)-2], s[index])
		}
	}

	for i := 0; i < numRows; i++ {
		for _, val := range str[i] {
			res += string(val)
		}
	}

	return res
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""
	str := make([]string, numRows)

	curRow := 0
	goDown := false
	index := 0

	for ; index < len(s); index++ {
		str[curRow] += string(s[index])
		if curRow == 0 || curRow == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	for _, value := range str {
		res += value
	}

	return res
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
