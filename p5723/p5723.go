package main

import "fmt"

func main() {
	var l int
	fmt.Scanf("%d", &l)
	nums := getPrimeNumber(l)
	i, tot := 0, 0
	for i < len(nums) && tot + nums[i] <= l {
		tot += nums[i]
		i++
	}

	for j := 0; j < i; j++ {
		fmt.Println(nums[j])
	}
	fmt.Println(i)
}

func getPrimeNumber(upperBound int) []int {
	var res []int
	isPrime := make([]bool, upperBound + 1)
	for i := 0; i < len(isPrime); i++ {
		isPrime[i] = true
	}

	for i := 2; i <= upperBound; i++ {
		if !isPrime[i] {
			continue
		}
		res = append(res, i)
		for j := i * 2; j <= upperBound; j += i {
			isPrime[j] = false
		}
	}

	return res
}