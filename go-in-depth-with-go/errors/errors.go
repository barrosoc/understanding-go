package main

import "fmt"

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}

func main() {
	amount, error := paintNeeded(4.2, -3.0)
	fmt.Println(error)
	fmt.Printf("%0.2f litres needed\n", amount)
}
