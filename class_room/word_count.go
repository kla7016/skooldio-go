package class_room

import "strings"

func wordCount(s string) map[string]int {
	listString := strings.Split(s, " ")
	mapOutput := make(map[string]int)

	for _, v := range listString {
		mapOutput[v] = mapOutput[v] + 1
	}
	return mapOutput
}
