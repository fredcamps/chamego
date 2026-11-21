package main

import (
	"fmt"

	"github.com/fredcamps/chamego/internal/chame"
)

func main() {
	g := chame.Chame("key", 1, false)
	fmt.Println(g)
}
