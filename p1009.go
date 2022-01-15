package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	
	res := "0"
	for i := 1; i <= n; i++ {
		res = addString(res, factorial(i))
	}

	fmt.Println(res)
}

func factorial(n int) string {
	if n <= 2 {
		return fmt.Sprintf("%d", n)
	}

	res, i := "1", n
	for i > 1 {
		res = multiplyString(res, fmt.Sprintf("%d", i))
		i--
	}

	return res
}

func addString(a, b string) string {
	aRvs, bRvs := reverseString(a), reverseString(b)
	arr := make([]byte, len(aRvs) + len(bRvs))
	minSize := minInt(len(aRvs), len(bRvs))
	for i := 0; i < minSize; i++ {
		addArray(arr, i, (aRvs[i] - '0') + (bRvs[i] - '0'))
	}

	rest := ""
	if minSize < len(a) {
		rest = aRvs
	}
	if minSize < len(b) {
		rest = bRvs
	}
	for i := minSize; i < len(rest); i++ {
		addArray(arr, i, rest[i] - '0')
	}

	return bytes2Str(arr)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func multiplyString(a, b string) string {
	aRvs, bRvs := reverseString(a), reverseString(b)
	arr := make([]byte, len(a)+len(b))
	for i := 0; i < len(aRvs); i++ {
		for j := 0; j < len(bRvs); j++ {
			multiplyRes := (aRvs[i] - '0') * (bRvs[j] - '0')
			addArray(arr, i+j, multiplyRes)
		}
	}

	return bytes2Str(arr)
}

func addArray(b []byte, pos int, num byte) {
	for pos < len(b) {
		b[pos] += num
		num = b[pos] / 10
		b[pos] %= 10
		pos++
	}
}

func reverseString(str string) string {
	res := ""
	for _, ch := range str {
		res = fmt.Sprintf("%c%s", ch, res)
	}

	return res
}

func bytes2Str(arr []byte) string {
	i := len(arr) - 1
	for i >= 0 && arr[i] == 0 {
		i--
	}
	if i < 0 {
		return "0"
	}

	res := ""
	for i >= 0 {
		res = fmt.Sprintf("%s%d", res, arr[i])
		i--
	}
	return res
}