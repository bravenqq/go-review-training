// Package main provides ...
package main

import "errors"

func splitArray(array []int, part int) ([][]int, error) {
	var res [][]int
	if part == 0 || array == nil {
		return nil, errors.New("part 不能为零")
	}
	if part == 1 {
		res = append(res, array)
		return res, nil
	}

	average := len(array) / part
	rem := len(array) % part
	lenArr := make([]int, part)
	for i := 0; i < part; i++ {
		if i < rem {
			lenArr[i] = average + 1
			continue
		}
		lenArr[i] = average
	}

	var index int
	res = make([][]int, len(array))
	for i := 0; i < part; i++ {
		if lenArr[i] > 0 {
			if i == part-1 {
				res[i] = array[index:]
				continue
			}
			index += lenArr[i]
			res[i] = array[index-lenArr[i] : index]
		}
	}

	return res, nil

}
