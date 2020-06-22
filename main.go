package main

import "fmt"

func Converter(s [][]int) []int {
	var newSlice []int
	var k, l int = 0, 0
	col := len(s)
	row := col

	for k < col && l < row {
		for i := l; i < row; i++ {
			newSlice = append(newSlice, s[k][i])
		}
		k++

		for i := k; i < col; i++ {
			newSlice = append(newSlice, s[i][row - 1])
		}
		row--

		if k < col {
			for i := row - 1; i >= l; i-- {
				newSlice = append(newSlice, s[col - 1][i])
			}
			col--;
		}

		if l < row {
			for i := col - 1; i >= k; i-- {
				newSlice = append(newSlice, s[i][l])
			}
			l++;
		}
	}

	return newSlice
}

func main() {
	fmt.Println(Converter([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(Converter([][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}))
}
