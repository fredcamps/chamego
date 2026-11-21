package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Initialize the client
	client := core.NewClient("Chamego Client", "1.0.0")

	// Example usage
	fmt.Println(client.Info())

	// Your client logic here
	_ = ctx
	_ = log
}
