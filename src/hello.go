package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)


func media(n1 float64, n2 float64) float64 {
	baseN1 := 0.4
	baseN2 := 0.6

	return (n1 * baseN1) + (n2 * baseN2)
}

func parseFloat64(text string) float64 {
	f, err := strconv.ParseFloat(text, 64)

	if err != nil {
		fmt.Println("Numero digitado é invalido")
	}

	return f
}

func buildInterface(text string) string {
	var input string

	fmt.Print(text)
	fmt.Scanln(&input)

	return input
}

func main() {
	n1 := parseFloat64(buildInterface("Digite a nota N1 : "))
	n2 := parseFloat64(buildInterface("Digite a nota N2 : "))

	fmt.Println("Media :", math.Round(media(n1, n2)*100)/100)

	fmt.Println("Time ", time.Now(), n1)
}
