package main

import (
	"fmt"
)

func main() {
	names := []string{"anna", "boris", "valentin"}
	var char string
	fmt.Scan(&char)
	result := getResult(names, char)
	fmt.Println(result)
}

func getResult (array[]string, char string) []string {
	resultArr := make([]string, 0)
	for i := range array {
		if string(array[i][0]) == string(char) {
			resultArr = append(resultArr, array[i])
		}
	}
	return resultArr
}