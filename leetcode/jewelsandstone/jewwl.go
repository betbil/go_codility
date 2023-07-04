package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	fmt.Println("jewels", jewels, "stones", stones)
	jewelcnt := 0
	var stonemap = make(map[rune]int)
	for _, v := range stones {
		stonemap[v]++
	}

	fmt.Println(stonemap)

	for _, v := range jewels {
		if cnt, ok := stonemap[v]; ok {
			jewelcnt += cnt
		}
	}
	fmt.Println("jewels", jewels, "stones", stones, "jewelcnt", jewelcnt)
	return jewelcnt
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
