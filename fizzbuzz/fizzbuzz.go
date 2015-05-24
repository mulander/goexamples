package fizzbuzz

import (
	"fmt"
	"log"
)

func Generate(count int) ([]string, error) {
	if count <= 0 {
		return nil, fmt.Errorf("fizzbuzz: Negative fizzbuzz count provided")
	}

	fizzbuzz := make([]string, count)

	var output string
	for i := 1; i <= count; i++ {
		switch {
		case i%15 == 0:
			output = "FizzBuzz"
		case i%3 == 0:
			output = "Fizz"
		case i%5 == 0:
			output = "Buzz"
		default:
			output = fmt.Sprintf("%d", i)
		}
		fizzbuzz[i-1] = output
	}

	return fizzbuzz, nil
}

func Print(count int) {
	fizzbuzz, err := Generate(count)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range fizzbuzz {
		fmt.Println(entry)
	}
}
