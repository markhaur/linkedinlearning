package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	// mocking algorithm processing time
	time.Sleep(20 * time.Millisecond)

	return Bid{
		"https://adsrus.com/19",
		0.05,
	}
}

var defaultBid = Bid{
	"https://adsrus.com/default",
	0.02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)

	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

func main() {
	// everytime we create a context, we create it from another context
	// there is a primary background context which is the parent of all contexts
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "https://http.cat/418"
	bid := findBid(ctx, url)
	fmt.Println(bid)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url = "https://http.cat/404"
	bid = findBid(ctx, url)
	fmt.Println(bid)
}
