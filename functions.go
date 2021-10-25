package main

import "strconv"

func Invert(matrix [][]string) [][]string {
	var i, j int
	row := len(matrix)
	result := make([][]string, 0)
	for i = 0; i < row; i++ {
		column := len(matrix[i])
		temp := make([]string, 0)
		for j = 0; j < column; j++ {
			value := matrix[j][i]
			temp = append(temp, value)
		}
		result = append(result, [][]string{temp}...)
	}
	return result
}
func Math(matrix [][]string, isSum bool) (string, error) {
	var i, j int
	row := len(matrix)
	var total int = 0
	if !isSum {
		total = 1
	}
	for i = 0; i < row; i++ {
		column := len(matrix[i])
		for j = 0; j < column; j++ {
			num, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return err.Error(), err
			}
			if isSum {
				total += num
				continue
			}
			total *= num
		}
	}
	return strconv.Itoa(total), nil
}
