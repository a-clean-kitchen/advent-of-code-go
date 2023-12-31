package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/a-clean-kitchen/advent-of-code-go/aoc_lib"
	"github.com/a-clean-kitchen/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) (builtInt int) {
	for _, line := range strings.Split(input, "\n") {
		val, err := strconv.Atoi(findFirstDigit(line) + findFirstDigit(aoc_lib.ReverseString(line)))
		if err != nil {
			panic(err)
		}
		builtInt += val
	}
	return builtInt
}

func findFirstDigit(line string) (digit string) {
	digits := strings.Split(line, "")
	fmt.Println(digits)
	for _, c := range digits {
		if _, err := strconv.Atoi(c); err == nil {
			digit = c
			break
		}
	}
	return digit
}

func parseIntName(chars []string) {}

func part2(input string) int {
	return 1
}
