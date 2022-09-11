package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `jsom:"amount"`
}

var data = `
{
	"user": "Scrooge McDuck",
	"type": "deposit",
	"amount": 123.4
}
`

func main() {
	rdr := strings.NewReader(data)

	// Decode request
	dec := json.NewDecoder(rdr)

	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %#v\n", req)

	// Create response
	prevBalance := 1_000_000.0 // Loaded from database
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// Encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}

}
