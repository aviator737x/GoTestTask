package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)


var dict map[string][61]int

func toNum(s string) int {
	res := 0
	pow := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i] - '0') * pow
		pow *= 10
	}
	return res
}

func main()  {
	data, err := ioutil.ReadFile("/Users/user/go/src/awesomeProject/main/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	FileStr := strings.Split(string(data), "\n")
	sort.Strings(FileStr)
	dict = make(map[string][61]int)
	CurrentGuard := ""
	state := 0
	for i := 0; i < len(FileStr); i++ {
		str := strings.Split(string(FileStr[i]), " ")
		if str[2] == "Guard" {
			GuardNumb := str[3][1:]
			CurrentGuard = GuardNumb
			state = 0
		}
		if str[2] == "falls" && state != 1{
			minute := toNum(str[1][3:5])
			field := dict[CurrentGuard]
			for i := minute; i < 60; i++ {
				field[i]++
				field[60]++
			}
			dict[CurrentGuard] = field
			state = 1
		}
		if str[2] == "wakes" && state != 2{
			minute := toNum(str[1][3:5])
			field := dict[CurrentGuard]
			for i := minute; i < 60; i++ {
				field[i]--
				field[60]--
			}
			dict[CurrentGuard] = field
			state = 2
		}
	}
	max := 0
	guard := ""
	for i := range dict {
		if max < dict[i][60] {
			max = dict[i][60]
			guard = i
		}
	}
	MaxMinute := 0
	result := 0
	for i := 0; i < 60; i++ {
		if MaxMinute < dict[guard][i] {
			MaxMinute = dict[guard][i]
			result = i
		}
	}
	fmt.Println("The answer to task#1 is:", result * toNum(guard))
	minute := 0
	max = 0
	for i := range dict {
		for j := 0; j < 60; j++ {
			if max < dict[i][j] {
				max = dict[i][j]
				guard = i
				minute = j
			}
		}
	}
	fmt.Println("The answer to task#2 is:", minute * toNum(guard))

}

