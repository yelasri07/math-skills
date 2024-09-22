package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	mathskills "mathskills/program"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please enter just name file data")
		return
	}

	if !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("erreur .txt")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("erreur")
	}

	content := bufio.NewScanner(file)

	Numbers := []int{}

	for content.Scan() {

		nb, err := strconv.Atoi(content.Text())
		if err != nil {
			fmt.Println("Not a number")
			continue
		}

		Numbers = append(Numbers, nb)
	}

	if err := content.Err(); err != nil {
		return
	}

	fmt.Println("Average: ", int(math.Round(mathskills.Average(Numbers))))
	fmt.Println("Median: ", int(math.Round(mathskills.Median(Numbers))))
	fmt.Println("Variance: ", int(math.Round(mathskills.Variance(Numbers))))
	fmt.Println("Standard Deviation: ", int(math.Round(mathskills.StandardDeviation(Numbers))))
}
