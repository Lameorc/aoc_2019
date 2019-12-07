package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// initialize copies contents from s into i casted as integers
func initialize(i []int, s []string) {
	for p, c := range s {
		i[p], _ = strconv.Atoi(c)
	}
}

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(f), ",")
	// convert to ints for easier work
	i := make([]int, len(s))
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			initialize(i, s)
			pos := 0
			i[1] = noun
			i[2] = verb
			for {
				if i[pos] == 99 {
					break
				} else if i[pos] == 1 {
					i[i[pos+3]] = i[i[pos+1]] + i[i[pos+2]]
				} else if i[pos] == 2 {
					i[i[pos+3]] = i[i[pos+1]] * i[i[pos+2]]
				} else {
					panic(fmt.Sprintf("Unexpected directive %d at pos: %d", i[pos], i))
				}
				pos += 4
			}
			if i[0] == 19690720 {
				fmt.Println(fmt.Sprintf("noun is `%d`; verb is `%d`, times 100 is `%d`", noun, verb, 100*noun+verb))
				break
			}
		}
	}
}
