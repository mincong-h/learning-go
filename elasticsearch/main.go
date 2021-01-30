package main

import (
	"log"
	"os"
)

func PrintUsage() {
	panic(`Usage:

    go run elasticsearch/*.go <sample>

For example:

    go run elasticsearch/*.go info
    go run elasticsearch/*.go get_indices`)
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("Missing sample name. Args: %s", os.Args)
		PrintUsage()
	}
	sample := os.Args[1]
	switch sample {
	case "info":
		Info()
		break
	case "get_indices":
		GetIndices()
		break
	default:
		log.Printf("Unknown sample name: %s", sample)
		PrintUsage()
	}
}
