package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	aLen, bLen := intLength(a), intLength(b)
	var palindromes []int
	for i := aLen; i <= bLen; i++ {
		palindromes = append(palindromes, genPalindromes(i)...)
	}

	var res []int
	for _, num := range palindromes {
		if a <= num && num <= b && isPrime(num) {
			res = append(res, num)
		}
	}

	for _, num := range res {
		fmt.Println(num)
	}
}

func genPalindromes(nBit int) []int{
	if nBit == 1 {
		var res []int
		for i := 1; i < 10; i++ {
			res = append(res, i)
		}
		return res
	}

	withMid := nBit % 2 == 1
	nPrefix := nBit / 2
	var res []int
	start, end := int(math.Pow10(nPrefix - 1)), int(math.Pow10(nPrefix))
	for i := start; i < end; i++ {
		postfix, l := reverseInt(i)
		postfixLen := int(math.Pow10((l)))
		if !withMid {
			res = append(res, (i * postfixLen) + postfix)
			continue
		}

		for mid := 0; mid < 10; mid++ {
			res = append(res, (i * 10 + mid) * postfixLen + postfix)
		}
	}

	return res
}

func intLength(n int) int {
	var res int
	for n > 0 {
		n /= 10
		res++
	}

	return res
}

func reverseInt(n int) (int, int) {
	var l int
	tempN := n
	for tempN % 10 == 0 {
		tempN /= 10
		l++
	}

	var res int
	for n > 0 {
		res = res * 10 + n % 10
		n /= 10
	}

	return res, l + intLength(res)
}

func isPrime(num int) bool {
	rootNum := math.Sqrt(float64(num))
	for i := 2; i <= int(rootNum); i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}