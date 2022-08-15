package main

import (
	"fmt"
)

var (
	answer = map[string]int{
		"III":     3,
		"LVIII":   58,
		"MCMXCIV": 1994,
		"DCXXI":   621,
	}
)

func main() {
	for input, output := range answer {
		if output != romanToInt(input) {
			panic(fmt.Sprintf(`input %s, should output %d, actual output %d`, input, output, romanToInt(input)))
		}
	}

	println(`通过`)
}

var (
	char2Value = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt(s string) int {
	var (
		last    = 0 // 上一个值
		times   = 0 // 单个值的重复次数
		current = 0 // 当前值
		value   = 0 // 完整值
	)

	for _, char := range s {
		current = char2Value[char]

		if last == 0 { // 重新一段新的，从头计数
			last = current
			times = 1

			continue
		}

		if last == current { // 重复,添加计数
			times++

			continue
		}

		if last > current { // 从大到小，开始累计
			value += last * times
			last = current
			times = 1
			continue
		}

		if (last == 1 || last == 10 || last == 100) && (current/last == 5 || current/last == 10) { // 从小到大
			value += current - last
			last, times = 0, 0
			continue
		}
	}

	value += last * times
	return value
}
