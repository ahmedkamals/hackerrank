package strings

import (
	"fmt"
	Strings "strings"
)

//func main() {
//
//	var n, k uint32
//	var str string
//	fmt.Scanf("%v", &n)
//
//	for k = 0; k < n; k++ {
//		fmt.Scanf("%v", &str)
//		fmt.Println(getCost(str))
//	}
//}


func (* StringUtil) GetCost(str string) int {

	var partialString, partialEndString string
	cost := 0

	for key, value := range str {

		partialString += string(value)
		partialEndString = str[key + 1 :]
		cost++

		fmt.Println(partialString, partialEndString)

		if Strings.Contains(partialString, partialEndString) {
			break
		}
	}

	return cost
}
