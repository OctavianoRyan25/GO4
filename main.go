package main

import "fmt"

func main() {
	word := "SELAMAT MALAM"
	var word2 = []string{"S", "E", "L", "A", "M", "A", "T", " ", "M", "A", "L", "A", "M"}
	for _, char := range word {
		fmt.Printf("%c\n", char)
	}

	dup_map := dup_count(word2)
	for k, v := range dup_map {
		fmt.Printf("%s : %d ", k, v)
	}

}

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1 
		} else {
			duplicate_frequency[item] = 1 
		}
	}
	return duplicate_frequency
}