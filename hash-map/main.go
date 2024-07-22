package main

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	sMap := make(map[rune]int)
	for _, v := range sentence {
		sMap[v]++
	}

}
