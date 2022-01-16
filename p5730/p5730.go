package main

import "fmt"

var bitMap = [][][]bool{
	{
		{true, true, true,},
		{true, false, true,},
		{true, false, true,},
		{true, false, true,},
		{true, true, true},
	},
	{
		{false, false, true,},
		{false, false, true,},
		{false, false, true,},
		{false, false, true,},
		{false, false, true,},
	},
	{
		{true, true, true,},
		{false, false, true,},
		{true, true, true},
		{true, false, false},
		{true, true, true},
	},
	{
		{true, true, true},
		{false, false, true},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	},
	{
		{true, false, true},
		{true, false, true},
		{true, true, true},
		{false, false, true},
		{false, false, true},
	},
	{
		{true, true, true},
		{true, false, false},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	},
	{
		{true, true, true},
		{true, false, false},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	},
	{
		{true, true, true},
		{false, false, true},
		{false, false, true},
		{false, false, true},
		{false, false, true},
	},
	{
		{true, true, true},
		{true, false, true},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	},
	{
		{true, true, true},
		{true, false, true},
		{true, true, true,},
		{false, false, true},
		{true, true, true},
	},
}

func main() {
	var b int
	var num string
	fmt.Scanf("%d\n%s", &b, &num)

	for i := 0; i < 5; i++ {
		for j := 0; j < b; j++ {
			shouldPrint := bitMap[num[j]-'0'][i]
			for k := 0; k < 3; k++ {
				if shouldPrint[k] {
					fmt.Print("X")
					continue
				}
				fmt.Print(".")
			}
			if j == b - 1 {
				continue
			}

			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}