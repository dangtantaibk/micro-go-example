package helper

import "fmt"

func CheckSumSHA256(params ...string) string {
	 result := ""
	for p, _ := range params {
		fmt.Sprintf("%v|", p)
	}
	// md5 chuoi result
	return result
}
