package array

import "sort"

// Notice: string swap, sort.Slice, hash slice
// 1. string CAN'T swap in-place
// 2. sort.Slice(sliceVar, func(i int, j int) bool) is in-place
// 3. string([]rune | []byte) are hash kind value

func GroupAnagrams(strs []string) [][]string {
	anagrams := make([][]string, 0)
	charCounter := make(map[string][]string)

	// way 1: sort string
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i int, j int) bool {
			return bytes[i] < bytes[j]
		})
		sortStr := string(bytes)
		charCounter[sortStr] = append(charCounter[sortStr], str)
	}

	// way 2: sort alphabetic byte slice
	// for _, str := range strs {
	//     hash := hashWord(str)
	//     charCounter[hash] = append(charCounter[hash], str)
	// }

	for _, v := range charCounter {
		anagrams = append(anagrams, v)
	}

	return anagrams
}

func HashWord(str string) string {
	x := make([]byte, 26)
	for _, c := range str {
		x[c-97] += 1
	}
	return string(x)
}
