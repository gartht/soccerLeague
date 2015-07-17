package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gartht/minDeltaUtil"
)

func main() {
	filepath := os.Args[1]

	file, error := os.Open(filepath)

	if error != nil {
		panic(error)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	deltaFinder := minDeltaUtil.Finder(6, 7, 1)

	teamName := ""

	//Read over the first line to get rid of the header
	reader.ReadString('\n')

	for {
		line, error := reader.ReadString('\n')

		if error == nil {
			//clean up non data rows
			line = strings.Replace(line, "-", "", -1)

			if len(line) > 4 {
				teamName = deltaFinder(strings.Fields(line))
			}

			continue
		}

		if error == io.EOF {
			break
		}

		panic(error)
	}
	fmt.Println(teamName)
}
