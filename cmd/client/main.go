package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fredcamps/chamego/pkg/core"
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
