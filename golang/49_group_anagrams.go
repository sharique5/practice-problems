package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string]string)
	for i := 0; i < len(strs); i++ {
		temp := strings.Split(strs[i], "")
		sort.Strings(temp)
		sorted := strings.Join(temp, "")
		if _, ok := groupMap[sorted]; ok {
			groupMap[sorted] += ", " + strs[i]
		} else {
			groupMap[sorted] = strs[i]
		}
	}
	res := make([][]string, len(groupMap))
	var ind int = 0
	for _, v := range groupMap {
		res[ind] = append(strings.Split(v, ", "))
		ind++
	}
	return res
}

func main() {
	var inp string
	fmt.Scanf("%s", &inp)
	fmt.Println("innnnn::: ", inp)
	inpList := strings.Split(inp, ", ")
	fmt.Println("inoutList:: ", inpList)
	fmt.Println(groupAnagrams(inpList))
}
