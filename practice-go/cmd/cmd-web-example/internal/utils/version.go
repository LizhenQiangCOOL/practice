package utils

import (
	"fmt"
	"os"
)

var version string

func getVersion() string {
	return version
}

func PrintVersion() {
	version := getVersion()
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Version", version)
}
