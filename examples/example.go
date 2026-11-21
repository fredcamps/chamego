package main

import (
	"fmt"

	"github.com/fredcamps/chamego/pkg/core"
	"github.com/fredcamps/chamego/pkg/utils"
)

func main() {
	p := core.New("chamego", "1.0.0")
	fmt.Println(p.Info())

	items := []string{"hello", "world"}
	fmt.Println(utils.Join(", ", items))
	fmt.Println(utils.Format("example", 5))
}
