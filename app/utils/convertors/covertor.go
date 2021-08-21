package convertors

import (
	"strconv"
	"strings"
)

func StringListOfStringsToIntList(l string) []int {
	slice := ParseStringsList(l)
	if len(slice) <= 0 || slice[0] == "" {
		return []int{}
	}
	res := StringSliceToInt(slice)

	return res
}

func ParseStringsList(l string) []string {
	if len(l) <= 0 {
		return nil
	}
	buffer := strings.Replace(l, "[", "", 1)
	buffer = strings.Replace(buffer, "]", "", 1)
	slice := strings.Split(buffer, ",")
	return slice
}

func StringSliceToInt(slice []string) []int {
	var res []int
	for _, e := range slice {
		intVal, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		} else {
			res = append(res, intVal)
		}
	}

	return res
}