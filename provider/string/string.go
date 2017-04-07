package string

import "fmt"

var Print = fmt.Println

//Get the partition of the string str
func Substr(str string, begin int, ends ...int) string {
	var ret string
	var end int = -1

	if len(ends) >= 0 {
		end = ends[0]
	}

	strSlic := []rune(str)

	strLen := len(strSlic)

	if end > strLen {
		end = strLen
	}

	end = end - begin

	ret = string(strSlic[begin:end])
	return ret
}
