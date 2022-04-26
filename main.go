package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/chart"
)

func main() {
	fmt.Println("Hello world!")
	c := chart.Chart{}
	fmt.Printf("a chart: %s\n", c)
}
