package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	ID   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "random123")
	f1(ctx)
	/* ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go getProductDetailsExternalService(1)
	select {
	case details := <-productChannel:
		fmt.Println("Details: ", details)
	case <-ctx.Done():
		fmt.Println("Timeout occurred, context cancelled!")
	}
	fmt.Println("End of the main") */

}

func f1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	f2(ctx)
}

func f2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	f3(ctx)
}

func f3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}

func getProductDetailsExternalService(productId int64) {
	time.Sleep(4 * time.Second)

	productChannel <- Product{
		ID:   productId,
		Name: "Kettle",
	}

}
