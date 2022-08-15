package main

import (
	"log"
	"strings"
)

var (
	answers = map[string][]string{
		"fl": []string{"flower", "flow", "flight"},
	}
)

func main() {
	for answer, input := range answers {
		result := longestCommonPrefix(input)
		if result != answer {
			log.Fatalf(`失败,预期输出%s,实际输出%s,输入%s`, answer, result, strings.Join(input, ","))
		}
	}
}

func longestCommonPrefix(candidates []string) string {
	shortestIndex := -1
	shortestLength := -1

	toMatch := make(map[string]struct{}, len(candidates))

	for i, str := range candidates { // 遍历，找到最短的，并且记录所有字符串
		toMatch[candidates[i]] = struct{}{}

		if shortestIndex == -1 || len(str) < shortestLength {
			shortestIndex = i
			shortestLength = len(str)
		}
	}

	delete(toMatch, candidates[shortestIndex]) // 去掉最短的本身

	for shortestLength >= 1 { // 保证有字符串
		prefix := candidates[shortestIndex][:shortestLength]

		for str, _ := range toMatch { // 遍历其他
			if strings.HasPrefix(str, prefix) { // 如果匹配，那就移除，因为更短的肯定匹配
				delete(toMatch, str)
			}
		}

		if len(toMatch) == 0 { // 没有了，就是找到了
			return prefix
		}

		shortestLength-- // 否则，继续缩短
	}

	return ""
}
