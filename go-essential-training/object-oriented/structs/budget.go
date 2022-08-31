package main

import (
	"fmt"
	"time"
)

// Budget is a budgest for campaign
/**
In go, it is very simple to specify private and public access modifiers
Everything which starts with Capital letter is exposed outside of package.
Everything which starts with small letter is not exposed outside of package.
In go, we call it Exported and Unexported symbols.
*/
type Budget struct {
	CampaignID string
	Balance    float64 // USD
	Expires    time.Time
}

// we can also define methods on structs
// the method also takes the receiver
func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

func NewBudget(campaignId string, balance float64, expires time.Time) (*Budget, error) {
	if campaignId == "" {
		return nil, fmt.Errorf("campaign id cannot be null")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be greater than 0")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		CampaignID: campaignId,
		Balance:    balance,
		Expires:    expires,
	}
	return &b, nil
}

func main() {
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)
	fmt.Println(b1.TimeLeft())

	b1.Update(10.5)
	fmt.Println(b1.Balance)

	// if you want to print more detail of struct Budget
	fmt.Printf("%#v\n", b1)

	// we can use . notation to access single value in a struct
	fmt.Println(b1.CampaignID)

	// we can also create Structs by specifying name, not positions
	b2 := Budget{
		Balance:    19.3,
		CampaignID: "puppies",
	}
	fmt.Printf("%#v\n", b2)

	// also we can specify the whole budget with the default values
	var b3 Budget
	fmt.Printf("%#v\n", b3)

	// creating budget with NewBudget function
	nb, _ := NewBudget("newkittens", 10.3, time.Now().Add(7*24*time.Hour))
	// if you want to print more detail of struct Budget
	fmt.Printf("%#v\n", nb)
}
