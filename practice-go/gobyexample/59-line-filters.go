package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	// echo 'hello' > /tmp/lines
	// echo 'filter' >> /tmp/line


	scanner := bufio.NewScanner(os.Stdin)


	for scanner.Scan(){
		url := strings.ToUpper(scanner.Text())

		fmt.Println(url)
	}

	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}