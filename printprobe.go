package main

import (
	"fmt"
	"github.com/sidepelican/goprobe/probe"
)

func main() {
	source, err := probe.FindAndNewProbeSource()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer source.Close()

	for record := range source.Records() {
		fmt.Println(record.String())
	}
}