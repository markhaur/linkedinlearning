package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 1,
		"GOOG": 2,
		"MSFT": 3, // must have trailing comma in the multiline
	}

	// prints the size of map
	fmt.Println(len(stocks))

	// get a value from map
	fmt.Println(stocks["AMZN"])

	// get zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	// use two values to see if found
	value, ok := stocks["TSLA"]

	if !ok {
		fmt.Println("not found!")
	} else {
		fmt.Println(value)
	}

	// set the value
	stocks["TSLA"] = 4
	fmt.Println(stocks["TSLA"])

	// Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is on keys and values
	for key, value := range stocks {
		fmt.Printf("%v -> %v\n", key, value)
	}

}
