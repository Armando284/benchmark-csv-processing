package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	testsNum, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("ls", "-l") // test command

	combined, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(combined))

	tests := [4]string{"node js/test1.js", "python python/test1.py", "php php/test1.php", "go run go/test1.go"}

	for _, test := range tests {
		for i := 0; i < int(testsNum); i++ {
			arr := strings.Fields(test)
			fmt.Println("Test: ", i+1, arr)
			command := arr[0]
			args := arr[1:]
			cmd := exec.Command(command, args...)

			_, err := cmd.CombinedOutput()
			if err != nil {
				panic(err)
			}

		}
	}

	fmt.Println("All", testsNum, "tests done. Results at ./output/benchmark.csv")
}
