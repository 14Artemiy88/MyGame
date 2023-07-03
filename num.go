package main

import (
	"strconv"
	"strings"
)

var bigNum = [][]string{
	{
		" #### ",
		"##  ##",
		"##  ##",
		"##  ##",
		" #### ",
	},
	{
		"  ##",
		"####",
		"  ##",
		"  ##",
		"  ##",
	},
	{
		" #### ",
		"##  ##",
		"   ## ",
		" ##   ",
		"######",
	},
	{
		" #### ",
		"#   ##",
		"  ### ",
		"#   ##",
		" #### ",
	},
	{
		"##    ",
		"##  ##",
		"######",
		"    ##",
		"    ##",
	},
	{
		"##### ",
		"##    ",
		"##### ",
		"    ##",
		"##### ",
	},
	{
		" #### ",
		"##    ",
		"##### ",
		"##  ##",
		" #### ",
	},
	{
		"######",
		"##  ##",
		"   ## ",
		"  ##  ",
		" ##   ",
	},
	{
		" #### ",
		"##  ##",
		" #### ",
		"##  ##",
		" #### ",
	},
	{
		" #### ",
		"##  ##",
		" #####",
		"    ##",
		" #### ",
	},
}

func getBigNum(num int) [][]string {
	numArray := strings.Split(strconv.Itoa(num), "")
	var newBigNum [][]string
	for _, i := range numArray {
		index, _ := strconv.Atoi(i)
		newBigNum = append(newBigNum, bigNum[index])
	}

	return newBigNum
}
