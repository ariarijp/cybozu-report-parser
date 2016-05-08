package main

import (
	"fmt"
	"os"

	"github.com/ariarijp/cybozu-report-parser"
)

func main() {
	schedules := cybozureport.Parse(os.Args[1])
	for _, s := range schedules {
		fmt.Println(s)
	}
}
