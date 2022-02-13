package main

import (
	"fmt"
	"github.com/mkmik/argsort"
	"github.com/odysseus/vigenere"
	"sort"
	"strings"
)

type Slice struct {
	sort.Float64Slice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(n ...float64) *Slice {
	s := &Slice{Float64Slice: sort.Float64Slice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)
}

func findCipherLength(input string) {
	input = removeCharacters(input, " .,?!\n")
	input = strings.ToUpper(input)

	fmt.Println(input)

	for i := 1; i < len(input); i++ {
		counter := 0
		for j := range input {
			index := i + j
			if index >= len(input) {
				index -= len(input)
			}

			if input[j] == input[index] {
				counter++
			}
		}
		fmt.Println(counter)
	}
}

func bruteForceText(input string, minLength int, maxLength int) {
	for i := minLength; i <= maxLength; i++ {
		counter := make([]rune, i)
		for j := 0; j < len(counter); j++ {
			counter[j] = 'a'
		}

		i := 0
		done := false
		for true {
			decoded := vigenere.Decipher(input, string(counter))

			//fmt.Printf("%s -> %s\n", string(counter), decoded)

			if strings.Contains(decoded, "SCHEICHEL") {
				fmt.Printf("\n%s -> %s\n", string(counter), decoded)
				done = true
				break
			}

			counter[0]++
			for j := 0; j < len(counter)-1; j++ {
				if counter[j] == 'z' {
					counter[j+1]++
					counter[j] = 'a'
				}
			}

			if i%100000 == 0 {
				fmt.Printf("%s -> %s\n", string(counter), decoded)
			}

			if counter[len(counter)-1] == 'z' {
				break
			}

			i++
		}

		if done {
			break
		}
	}
}

func sloveText(input string, length int, deepness int) {
	input = removeCharacters(input, " .,?!\n")
	input = strings.ToUpper(input)

	germanFrequenz := []float64{
		0.0558,
		0.0196,
		0.0316,
		0.0498,
		0.1693,
		0.0149,
		0.0302,
		0.0498,
		0.0802,
		0.0024,
		0.0132,
		0.0360,
		0.0255,
		0.1053,
		0.0224,
		0.0067,
		0.0002,
		0.0689,
		0.0642,
		0.0579,
		0.0383,
		0.0084,
		0.0178,
		0.0005,
		0.0005,
		0.0121,
	}

	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	abcLookUp := map[rune]int{}
	for i, r := range abc {
		abcLookUp[r] = i
	}

	results := make([][]int, length)
	for i := 0; i < length; i++ {
		var counter float64
		counterList := make([]int, len(abc))

		for j := i; j < len(input); j += length {
			r := rune(input[j])
			counterList[abcLookUp[r]]++
			counter++
		}

		frequenzList := make([]float64, len(abc))
		for j, n := range counterList {
			frequenzList[j] = float64(n) / counter
		}

		sums := make([]float64, len(abc))
		for j := 0; j < len(abc); j++ {

			for k := 0; k < len(abc); k++ {

				index := k + j
				if index >= len(abc) {
					index -= len(abc)
				}
				sums[j] += frequenzList[index] * germanFrequenz[k]
			}
		}

		indexs := argsort.Sort(sort.Float64Slice(sums))
		result := make([]int, len(abc))
		for j, n := range indexs {
			result[len(abc)-1-j] = n
		}

		results[i] = result
	}

	counter := make([]int, length)

	i := 0
	for true {

		key := ""
		for j, n := range counter {
			key += string(abc[results[j][n]])
		}

		decoded := vigenere.Decipher(input, key)

		if strings.Contains(decoded, "SCHEICHEL") {

			fmt.Printf("%d %s -> %s\n", counter, key, decoded)
			break
		}

		counter[0]++
		for j := 0; j < len(counter)-1; j++ {
			if counter[j] >= deepness {
				counter[j+1]++
				counter[j] = 0
			}
		}

		if i%1000 == 0 {
			fmt.Printf("%d %s -> %s\n", counter, key, decoded)
		}

		if counter[len(counter)-1] == deepness {
			break
		}

		i++
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
